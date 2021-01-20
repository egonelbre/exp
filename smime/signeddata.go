package smime

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"

	"github.com/egonelbre/exp/smime/ber"
)

const panictodo = false

type SignedData struct {
	Version          int
	DigestAlgorithms []pkix.AlgorithmIdentifier

	ContentType asn1.ObjectIdentifier
	Content     []byte

	Certificates []*x509.Certificate
	CRLs         []*pkix.CertificateList

	Signers []*Signer
}

func (s *SignedData) findCertFor(signer *Signer) (*x509.Certificate, error) {
	for _, cert := range s.Certificates {
		if signer.Id.Matches(cert) {
			return cert, nil
		}
	}
	return nil, fmt.Errorf("did not find certificate for signer %v", s)
}

func (s *SignedData) CheckSignature() (err error) {
	if len(s.Signers) == 0 {
		return fmt.Errorf("no signature found")
	}

	for _, signer := range s.Signers {
		var cert *x509.Certificate
		if cert, err = s.findCertFor(signer); err != nil {
			return
		}

		// check the revocation lists
		revoked := false
		for _, crl := range s.CRLs {
			if revokeerr := cert.CheckCRLSignature(crl); revokeerr != nil {
				revoked = true
			}
		}

		if revoked {
			if panictodo {
				panic("TODO: check counter signature")
			}
			// return fmt.Errorf("smime: certificate has been revoked")
		}

		data := s.Content
		if signer.RawSignedAttributes != nil {
			var attrvalue []byte
			if attrvalue, err = findSingleAttribute(signer.SignedAttributes, oidMessageDigest); err != nil {
				return errors.New("smime: message digest not found")
			}

			var digest []byte

			// the single message digest attribute should be OCTET STRING
			if _, err = asn1.Unmarshal(attrvalue, &digest); err != nil {
				return
			}

			calculated := data
			if hash := getDigestHashType(signer.DigestAlgorithm.Algorithm); hash.Available() {
				h := hash.New()
				h.Write(data)
				calculated = h.Sum(nil)
			}

			if !bytes.Equal(digest, calculated) {
				return errors.New("smime: message digest does not match content")
			}

			data = signer.RawSignedAttributes
		}

		algo := getSignatureAlgorithmFromOID(signer.SignatureAlgorithm.Algorithm)
		err = cert.CheckSignature(algo, data, signer.Signature)
		if err != nil {
			return
		}
	}
	return
}

type Signer struct {
	Id SignerId

	RawSignedAttributes []byte
	SignedAttributes    []Attribute
	UnsignedAttributes  []Attribute

	DigestAlgorithm    pkix.AlgorithmIdentifier
	SignatureAlgorithm pkix.AlgorithmIdentifier

	Signature []byte
}

// Either IssuerAndSerialNumber or SubjectKeyIdentifier
type SignerId interface {
	Matches(*x509.Certificate) bool
}

type IssuerAndSerialNumber struct {
	Issuer       pkix.Name
	SerialNumber *big.Int
}

func (sid IssuerAndSerialNumber) Matches(cert *x509.Certificate) bool {
	if panictodo {
		panic("issuer comparison")
	}
	return (sid.Issuer.CommonName == cert.Issuer.CommonName) &&
		(sid.SerialNumber.Cmp(cert.SerialNumber) == 0)
}

type SubjectKeyIdentifier []byte

func (sid SubjectKeyIdentifier) Matches(cert *x509.Certificate) bool {
	return bytes.Equal(cert.SubjectKeyId, sid)
}

type Attribute struct {
	Type   asn1.ObjectIdentifier
	Values [][]byte
}

func findSingleAttribute(attributes []Attribute, oid asn1.ObjectIdentifier) ([]byte, error) {
	for _, attr := range attributes {
		if attr.Type.Equal(oid) {
			if len(attr.Values) != 1 {
				return nil, errors.New("smime: attribute had improper amount of values")
			}
			return attr.Values[0], nil
		}
	}
	return nil, errors.New("smime: did not find required attribute")
}

func convertAttributeSet(attrs asnAttributeSet) (ret []Attribute) {
	for _, attr := range attrs.Attrs {
		a := Attribute{attr.Type, nil}
		for _, val := range attr.Values {
			data, _ := ber.EncodeTree(val)
			a.Values = append(a.Values, data)
		}
		ret = append(ret, a)
	}
	return
}

/*
	SignedData ::= SEQUENCE {
		version CMSVersion,
		digestAlgorithms DigestAlgorithmIdentifiers,
		encapContentInfo EncapsulatedContentInfo,
		certificates [0] IMPLICIT CertificateSet OPTIONAL,
		crls [1] IMPLICIT RevocationInfoChoices OPTIONAL,
		signerInfos SignerInfos }

	SignerInfos ::= SET OF SignerInfo
*/
type asnSignedData struct {
	Version      int64
	Algorithms   asnDigestAlgorithms
	EncapContent asnEncapsulatedContentInfo
	Certificates asnCertificateSet
	Crls         asnRevocationInfoChoices
	SignerInfos  asnSignerInfos
}

type crl struct{}

func (d *asnSignedData) marshaler() ber.Marshaler {
	return ber.Sequence{
		ber.Check{ber.Universal, ber.TagInteger, ber.Int64{&d.Version}},
		ber.Check{ber.Universal, ber.TagSet, &d.Algorithms},
		ber.Check{ber.Universal, ber.TagSequence, &d.EncapContent},
		ber.Optional{ber.Check{ber.Context, 0, &d.Certificates}},
		ber.Optional{ber.Check{ber.Context, 1, &d.Crls}},
		ber.Check{ber.Universal, ber.TagSet, &d.SignerInfos},
	}
}

func (d *asnSignedData) Unmarshal(tree *ber.Tree) error {
	return d.marshaler().Unmarshal(tree)
}

func (d *asnSignedData) Marshal() (*ber.Tree, error) {
	return d.marshaler().Marshal()
}

func parseSignedData(tree *ber.Tree) (*SignedData, error) {
	var data asnSignedData
	err := data.Unmarshal(tree)
	if err != nil {
		return nil, err
	}

	signeddata := &SignedData{
		Version:          int(data.Version),
		DigestAlgorithms: data.Algorithms,

		ContentType: data.EncapContent.ContentType,
		Content:     data.EncapContent.Data,

		Certificates: data.Certificates,
		CRLs:         data.Crls,
	}

	for _, si := range data.SignerInfos {
		var sid SignerId
		if si.Issuer.SerialNumber != nil {
			sid = IssuerAndSerialNumber{si.Issuer.Issuer, si.Issuer.SerialNumber}
		} else {
			sid = SubjectKeyIdentifier(si.SubjectKeyId)
		}

		signer := &Signer{
			Id: sid,

			RawSignedAttributes: si.SignedAttrs.Raw,
			SignedAttributes:    convertAttributeSet(si.SignedAttrs),
			UnsignedAttributes:  convertAttributeSet(si.UnsignedAttrs),

			DigestAlgorithm:    si.DigestAlgorithm,
			SignatureAlgorithm: si.SignatureAlgorithm,
			Signature:          si.Signature,
		}
		signeddata.Signers = append(signeddata.Signers, signer)
	}

	return signeddata, nil
}

/*
	EncapsulatedContentInfo ::= SEQUENCE {
		eContentType ContentType,
		eContent [0] EXPLICIT OCTET STRING OPTIONAL }

	ContentType ::= OBJECT IDENTIFIER
*/
type asnEncapsulatedContentInfo struct {
	ContentType asn1.ObjectIdentifier
	Data        []byte
}

func (d *asnEncapsulatedContentInfo) marshaler() ber.Marshaler {
	return ber.Sequence{
		ber.Check{ber.Universal, ber.TagObjectIdentifier, ber.ObjectIdentifier{&d.ContentType}},
		ber.Optional{ber.Check{ber.Context, 0,
			ber.Explicit{
				ber.Check{ber.Universal, ber.TagOctetString, ber.OctetString{&d.Data}},
			},
		}},
	}
}

func (d *asnEncapsulatedContentInfo) Unmarshal(tree *ber.Tree) error {
	return d.marshaler().Unmarshal(tree)
}

func (d *asnEncapsulatedContentInfo) Marshal() (*ber.Tree, error) {
	return d.marshaler().Marshal()
}

/*
	DigestAlgorithmIdentifiers ::= SET OF DigestAlgorithmIdentifier
	DigestAlgorithmIdentifier ::= AlgorithmIdentifier
*/
type asnDigestAlgorithms []pkix.AlgorithmIdentifier

func (d *asnDigestAlgorithms) Unmarshal(tree *ber.Tree) (err error) {
	algorithms := make([]pkix.AlgorithmIdentifier, len(tree.Children))

	for i, child := range tree.Children {
		id := (*asnAlgorithmIdentifier)(&algorithms[i])
		if err = id.Unmarshal(child); err != nil {
			return
		}
	}
	*d = algorithms
	return
}

func (d *asnDigestAlgorithms) Marshal() (tree *ber.Tree, err error) {
	if tree, err = ber.ASN1Tree(d); err != nil {
		return
	}
	tree.Tag = ber.TagSet
	return tree, err
}

/*
   AlgorithmIdentifier ::= SEQUENCE {
        algorithm          OBJECT IDENTIFIER,
        parameters         ANY
   }
*/
type asnAlgorithmIdentifier pkix.AlgorithmIdentifier

func (d *asnAlgorithmIdentifier) Unmarshal(tree *ber.Tree) error {
	return ber.Sequence{
		ber.Check{ber.Universal, ber.TagObjectIdentifier, ber.ObjectIdentifier{&d.Algorithm}},
		ber.Optional{ber.ASN1RawValue{&d.Parameters}},
	}.Unmarshal(tree)
}

func (d *asnAlgorithmIdentifier) Marshal() (*ber.Tree, error) {
	panic("unimplemented")
	return nil, nil
}

/*
CertificateSet	::=	 SET OF CertificateChoices
CertificateChoices	::=	CHOICE {
	certificate	Certificate,
	extendedCertificate	[0] IMPLICIT ExtendedCertificate, -- Obsolete
	v1AttrCert	 	[1] IMPLICIT AttributeCertificateV1, -- Obsolete
	v2AttrCert	 	[2] IMPLICIT AttributeCertificateV2,
	other	 		[3] IMPLICIT OtherCertificateFormat
}
AttributeCertificateV2	::=	 SET OF AttributeCertificate
OtherCertificateFormat	::=	SEQUENCE {
	otherCertFormat	 OBJECT IDENTIFIER,
	otherCert	 ANY DEFINED BY otherCertFormat
}
*/
type asnCertificateSet []*x509.Certificate

func (d *asnCertificateSet) Unmarshal(tree *ber.Tree) (err error) {
	var cert *x509.Certificate
	var bytes []byte

	ret := make([]*x509.Certificate, 0, len(tree.Children))

	for _, child := range tree.Children {
		if bytes, err = ber.EncodeTree(child); err != nil {
			return
		}
		if cert, err = x509.ParseCertificate(bytes); err != nil {
			return
		}
		ret = append(ret, cert)
	}
	*d = ret
	return
}

func (d *asnCertificateSet) Marshal() (tree *ber.Tree, err error) {
	tree = &ber.Tree{
		Token: &ber.Token{
			Kind:  ber.Constructed,
			Class: ber.Universal,
			Tag:   ber.TagSet,
		},
	}

	var sub *ber.Tree
	for _, cert := range *d {
		if sub, err = ber.DecodeTree(cert.Raw); err != nil {
			return nil, err
		}
		tree.Children = append(tree.Children, sub)
	}
	return
}

/*
   RevocationInfoChoices ::= SET OF RevocationInfoChoice

   RevocationInfoChoice ::= CHOICE {
     crl CertificateList,
     other [1] IMPLICIT OtherRevocationInfoFormat }

   OtherRevocationInfoFormat ::= SEQUENCE {
     otherRevInfoFormat OBJECT IDENTIFIER,
     otherRevInfo ANY DEFINED BY otherRevInfoFormat }
*/
type asnRevocationInfoChoices []*pkix.CertificateList

func (d *asnRevocationInfoChoices) Unmarshal(tree *ber.Tree) (err error) {
	var crl *pkix.CertificateList
	var bytes []byte

	ret := make([]*pkix.CertificateList, 0, len(tree.Children))
	for _, child := range tree.Children {
		if bytes, err = ber.EncodeTree(child); err != nil {
			return
		}

		if crl, err = x509.ParseCRL(bytes); err != nil {
			return
		}
		ret = append(ret, crl)
	}
	*d = ret
	return
}

func (d *asnRevocationInfoChoices) Marshal() (tree *ber.Tree, err error) {
	return ber.ASN1Tree(*d)
}

/*
  SignerInfos ::= SET OF SignerInfo
*/
type asnSignerInfos []*asnSignerInfo

func (d *asnSignerInfos) Unmarshal(tree *ber.Tree) (err error) {
	ret := make([]*asnSignerInfo, len(tree.Children))
	for i, child := range tree.Children {
		v := &asnSignerInfo{}
		if err = v.Unmarshal(child); err != nil {
			return
		}
		ret[i] = v
	}
	*d = ret
	return
}

func (d *asnSignerInfos) Marshal() (tree *ber.Tree, err error) {
	tree = &ber.Tree{
		Token: &ber.Token{
			Kind:  ber.Constructed,
			Class: ber.Universal,
			Tag:   ber.TagSet,
		},
	}

	var sub *ber.Tree
	for _, info := range *d {
		if sub, err = info.Marshal(); err != nil {
			return
		}
		tree.Children = append(tree.Children, sub)
	}

	return
}

/*

  SignerInfo ::= SEQUENCE {
     version CMSVersion,
     sid SignerIdentifier,
     digestAlgorithm DigestAlgorithmIdentifier,
     signedAttrs [0] IMPLICIT SignedAttributes OPTIONAL,
     signatureAlgorithm SignatureAlgorithmIdentifier,
     signature SignatureValue,
     unsignedAttrs [1] IMPLICIT UnsignedAttributes OPTIONAL }

   SignerIdentifier ::= CHOICE {
     issuerAndSerialNumber IssuerAndSerialNumber,
     subjectKeyIdentifier [0] SubjectKeyIdentifier }

   SignatureValue ::= OCTET STRING
*/
type asnSignerInfo struct {
	Version            int64
	Issuer             asnIssuerAndSerialNumber
	SubjectKeyId       []byte
	DigestAlgorithm    pkix.AlgorithmIdentifier
	SignedAttrs        asnAttributeSet
	SignatureAlgorithm pkix.AlgorithmIdentifier
	Signature          []byte
	UnsignedAttrs      asnAttributeSet
}

func (d *asnSignerInfo) marshaler() ber.Marshaler {
	return ber.Sequence{
		ber.Check{ber.Universal, ber.TagInteger, ber.Int64{&d.Version}},
		ber.Choice{
			ber.Check{ber.Universal, ber.TagSequence, &d.Issuer},
			ber.Check{ber.Context, 0, ber.OctetString{&d.SubjectKeyId}},
		},
		ber.Check{ber.Universal, ber.TagSequence, (*asnAlgorithmIdentifier)(&d.DigestAlgorithm)},
		ber.Optional{ber.Check{ber.Context, 0, &d.SignedAttrs}},
		ber.Check{ber.Universal, ber.TagSequence, (*asnAlgorithmIdentifier)(&d.SignatureAlgorithm)},
		ber.Check{ber.Universal, ber.TagOctetString, ber.OctetString{&d.Signature}},
		ber.Optional{ber.Check{ber.Context, 1, &d.UnsignedAttrs}},
	}
}

func (d *asnSignerInfo) Unmarshal(tree *ber.Tree) error {
	return d.marshaler().Unmarshal(tree)
}

func (d *asnSignerInfo) Marshal() (*ber.Tree, error) {
	return d.marshaler().Marshal()
}

/*
  IssuerAndSerialNumber ::= SEQUENCE {
    issuer Name,
    serialNumber CertificateSerialNumber }

  CertificateSerialNumber ::= INTEGER
*/
type asnIssuerAndSerialNumber struct {
	Issuer       pkix.Name
	SerialNumber *big.Int
}

func (d *asnIssuerAndSerialNumber) marshaler() ber.Marshaler {
	return ber.Sequence{
		ber.Check{ber.Universal, ber.TagSequence, (*asnName)(&d.Issuer)},
		ber.Check{ber.Universal, ber.TagInteger, ber.BigInt{&d.SerialNumber}},
	}
}

func (d *asnIssuerAndSerialNumber) Unmarshal(tree *ber.Tree) (err error) {
	return d.marshaler().Unmarshal(tree)
}

func (d *asnIssuerAndSerialNumber) Marshal() (*ber.Tree, error) {
	return d.marshaler().Marshal()
}

type asnName pkix.Name

func (d *asnName) Unmarshal(tree *ber.Tree) (err error) {
	var bytes []byte
	if bytes, err = ber.EncodeTree(tree); err != nil {
		return
	}

	var rdns pkix.RDNSequence
	if _, err = asn1.Unmarshal(bytes, &rdns); err != nil {
		return
	}
	(*pkix.Name)(d).FillFromRDNSequence(&rdns)

	return
}

func (d *asnName) Marshal() (*ber.Tree, error) {
	panic("unimplemented")
	return nil, nil
}

type asnAttributeSet struct {
	Raw   []byte
	Attrs []asnAttribute
}

func (d *asnAttributeSet) Unmarshal(tree *ber.Tree) (err error) {
	ret := make([]asnAttribute, len(tree.Children))
	for i, child := range tree.Children {
		err = (&ret[i]).Unmarshal(child)
		if err != nil {
			return
		}
	}

	d.Attrs = ret
	temp := &ber.Tree{
		Token: &ber.Token{
			Kind:  ber.Constructed,
			Class: ber.Universal,
			Tag:   ber.TagSet,
		},
		Children: tree.Children,
	}
	d.Raw, err = ber.EncodeTree(temp)
	return
}

func (d *asnAttributeSet) Marshal() (*ber.Tree, error) {
	panic("unimplemented")
	return nil, nil
}

type asnAttribute struct {
	Type   asn1.ObjectIdentifier
	Values []*ber.Tree
}

func (d *asnAttribute) Unmarshal(tree *ber.Tree) (err error) {
	var temp *ber.Tree

	err = ber.Sequence{
		ber.Check{ber.Universal, ber.TagObjectIdentifier, ber.ObjectIdentifier{&d.Type}},
		ber.Check{ber.Universal, ber.TagSet, ber.RawTree{&temp}},
	}.Unmarshal(tree)

	if err != nil {
		return
	}

	d.Values = temp.Children
	return
}

func (d *asnAttribute) Marshal() (*ber.Tree, error) {
	panic("unimplemented")
	return nil, nil
}
