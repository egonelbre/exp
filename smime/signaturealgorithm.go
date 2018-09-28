package smime

import (
	"crypto"
	"crypto/x509"
	"encoding/asn1"
)

var (
	oidSignatureRSA             = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	oidSignatureMD2WithRSA      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 2}
	oidSignatureMD5WithRSA      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 4}
	oidSignatureSHA1WithRSA     = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 5}
	oidSignatureSHA256WithRSA   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 11}
	oidSignatureSHA384WithRSA   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 12}
	oidSignatureSHA512WithRSA   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 13}
	oidSignatureDSAWithSHA1     = asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 3}
	oidSignatureDSAWithSHA256   = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 4, 3, 2}
	oidSignatureECDSAWithSHA1   = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 1}
	oidSignatureECDSAWithSHA256 = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 2}
	oidSignatureECDSAWithSHA384 = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 3}
	oidSignatureECDSAWithSHA512 = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 4}

	//TODO: add the algorithms in RFC 3370 & RFC 5754
	oidDigestSHA1 = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 26}
)

/*
type SignatureAlgorithm int

func (s SignatureAlgorithm) hashType() (crypto.Hash, error) {
	for _, details := range signatureAlgorithmDetails {
		if s == details.algo {
			return details.hash
		}
	}
	return crypto.Hash(0), errors.New("unknown signature algorithm")
}
*/

var signatureAlgorithmDetails = []struct {
	algo x509.SignatureAlgorithm
	oid  asn1.ObjectIdentifier
	hash crypto.Hash
}{
	// {x509.PlainRSA, oidSignatureRSA, x509.RSA, crypto.Hash(0)},
	{x509.MD2WithRSA, oidSignatureMD2WithRSA, crypto.Hash(0)},
	{x509.MD5WithRSA, oidSignatureMD5WithRSA, crypto.MD5},
	{x509.SHA1WithRSA, oidSignatureSHA1WithRSA, crypto.SHA1},
	{x509.SHA256WithRSA, oidSignatureSHA256WithRSA, crypto.SHA256},
	{x509.SHA384WithRSA, oidSignatureSHA384WithRSA, crypto.SHA384},
	{x509.SHA512WithRSA, oidSignatureSHA512WithRSA, crypto.SHA512},
	{x509.DSAWithSHA1, oidSignatureDSAWithSHA1, crypto.SHA1},
	{x509.DSAWithSHA256, oidSignatureDSAWithSHA256, crypto.SHA256},
	{x509.ECDSAWithSHA1, oidSignatureECDSAWithSHA1, crypto.SHA1},
	{x509.ECDSAWithSHA256, oidSignatureECDSAWithSHA256, crypto.SHA256},
	{x509.ECDSAWithSHA384, oidSignatureECDSAWithSHA384, crypto.SHA384},
	{x509.ECDSAWithSHA512, oidSignatureECDSAWithSHA512, crypto.SHA512},
}

func getSignatureAlgorithmFromOID(oid asn1.ObjectIdentifier) x509.SignatureAlgorithm {
	for _, details := range signatureAlgorithmDetails {
		if oid.Equal(details.oid) {
			return details.algo
		}
	}
	return x509.UnknownSignatureAlgorithm
}

func getDigestHashType(oid asn1.ObjectIdentifier) crypto.Hash {
	//TODO: properly select digest

	for _, details := range signatureAlgorithmDetails {
		if oid.Equal(details.oid) {
			return details.hash
		}
	}

	if oid.Equal(oidDigestSHA1) {
		return crypto.SHA1
	}

	return crypto.Hash(0)
}
