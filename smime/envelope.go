package smime

/*
type EnvelopedData struct {
	Version int

	OriginatorCerts []interface{}
	OriginatorCRLs  []*pkix.CertificateList

	RecipientInfos []interface{}

	EncryptedContent    []byte
	EncryptionAlgorithm EncryptionAlgorithm
}

type KeyTransfer struct {
	RAWRid []byte

	Issuer               pkix.Name
	IssuerSerialNumber   *big.Int
	SubjectKeyIdentifier []byte

	EncryptionAlgorithm EncryptionAlgorithm
	EncryptedKey        []byte
}

type KeyAgreement struct {
}

// RFC 5652 6.1

// EnvelopedData ::= SEQUENCE {
//   version CMSVersion,
//   originatorInfo [0] IMPLICIT OriginatorInfo OPTIONAL,
//   recipientInfos RecipientInfos,
//   encryptedContentInfo EncryptedContentInfo,
//   unprotectedAttrs [1] IMPLICIT UnprotectedAttributes OPTIONAL }
//
// RecipientInfos ::= SET SIZE (1..MAX) OF RecipientInfo
type envelopedData struct {
	Version int

	OriginatorInfo originatorInfo  `asn1:"optional,tag:0"`
	RecipientInfos []asn1.RawValue `asn1:"set"`

	EncryptedContentInfo encryptedContentInfo
	UnprotectedAttrs     []pkix.AttributeTypeAndValue `asn1:"optional,set,tag:1"`
}

// OriginatorInfo ::= SEQUENCE {
//   certs [0] IMPLICIT CertificateSet OPTIONAL,
//   crls [1] IMPLICIT RevocationInfoChoices OPTIONAL }
type originatorInfo struct {
	Certs []asn1.RawValue `asn:"optional,set,tag:0"`
	Crls  []asn1.RawValue `asn:"optional,set,tag:1"`
}

// EncryptedContentInfo ::= SEQUENCE {
//   contentType ContentType,
//   contentEncryptionAlgorithm ContentEncryptionAlgorithmIdentifier,
//   encryptedContent [0] IMPLICIT EncryptedContent OPTIONAL }
type encryptedContentInfo struct {
	ContentType                asn1.ObjectIdentifier
	ContentEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedContent           asn1.RawValue `asn:"optional,tag:0"`
}

func parseEnvelopedData(data []byte) (interface{}, error) {
	env := envelopedData{}
	_, err := asn1.Unmarshal(data, &env)
	if err != nil {
		return nil, err
	}

	envelope := &EnvelopedData{}
	envelope.Version = env.Version

	for _, v := range env.OriginatorInfo.Certs {
		cert, err := parseCertificateChoice(v.FullBytes)
		if err != nil {
			return nil, err
		}
		envelope.OriginatorCerts = append(envelope.OriginatorCerts, cert)
	}

	for _, v := range env.OriginatorInfo.Crls {
		crl, err := x509.ParseCRL(v.FullBytes)
		if err != nil {
			return nil, err
		}
		envelope.OriginatorCRLs = append(envelope.OriginatorCRLs, crl)
	}

	for _, info := range env.RecipientInfos {
		rinfo, err := parseRecipientInfo(info)
		if err != nil {
			return nil, err
		}

		envelope.RecipientInfos = append(envelope.RecipientInfos, rinfo)

	}

	envelope.EncryptionAlgorithm = getEncryptionAlgorithmFromOID(env.EncryptedContentInfo.ContentEncryptionAlgorithm.Algorithm)
	envelope.EncryptedContent = env.EncryptedContentInfo.EncryptedContent.Bytes

	return envelope, nil
}
*/
