package smime

import (
	"bytes"
	"encoding/asn1"
	"fmt"

	"github.com/egonelbre/smime/ber"
)

// RFC 5652 12.1
// -- Content Type Object Identifiers
//
// id-ct-contentInfo OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs9(9) smime(16) ct(1) 6 }
//
// id-data OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs7(7) 1 }
//
// id-signedData OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs7(7) 2 }
//
// id-envelopedData OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs7(7) 3 }
//
// id-digestedData OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs7(7) 5 }
//
// id-encryptedData OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs7(7) 6 }
//
// id-ct-authData OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//     us(840) rsadsi(113549) pkcs(1) pkcs-9(9) smime(16) ct(1) 2 }
var (
	oidContentInfo = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 1, 6}

	oidData          = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}
	oidSignedData    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 2}
	oidEnvelopedData = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 3}
	oidDigestedData  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 5}
	oidEncryptedData = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 6}
	oidAuthData      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 16, 1, 2}
)

type Data []byte

func Parse(data []byte) (ret interface{}, err error) {
	d := ber.NewDecoder(bytes.NewReader(data))
	tree, err := ber.ParseTree(d)
	if err != nil {
		return nil, err
	}

	//pp(tree)

	// RFC 5652, 3
	// ContentInfo ::= SEQUENCE {
	//    contentType ContentType,
	//    content [0] EXPLICIT ANY DEFINED BY contentType }
	//
	// ContentType ::= OBJECT IDENTIFIER

	var contentType asn1.ObjectIdentifier
	var content *ber.Tree

	err = ber.Sequence{
		ber.Check{ber.Universal, ber.TagObjectIdentifier, ber.ObjectIdentifier{&contentType}},
		ber.Check{ber.Context, 0, ber.Explicit{ber.RawTree{&content}}},
	}.Unmarshal(tree)
	if err != nil {
		return
	}

	switch {
	case oidData.Equal(contentType):
		bytes, err := content.AsOctetString()
		return Data(bytes), err
	case oidSignedData.Equal(contentType):
		ret, err := parseSignedData(content)
		return ret, err
	case oidEnvelopedData.Equal(contentType):
		return nil, fmt.Errorf("enveloped-data not suppoerted")
	case oidDigestedData.Equal(contentType):
		return nil, fmt.Errorf("digested-data not supported")
	case oidEncryptedData.Equal(contentType):
		return nil, fmt.Errorf("encrypted-data not supported")
	case oidAuthData.Equal(contentType):
		return nil, fmt.Errorf("auth-data not supported")
	}
	return nil, fmt.Errorf("unknown content type")
}
