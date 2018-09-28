package smime

import "encoding/asn1"

type EncryptionAlgorithm int

const (
	UnknownEncryptionAlgorithm EncryptionAlgorithm = iota
	RSA
	RC2_CBC
	RC4
	DES_EDE3_CBC
	RC5_CBC_Pad
	ID_DES_CDMF
)

func getEncryptionAlgorithmFromOID(oid asn1.ObjectIdentifier) EncryptionAlgorithm {
	for _, details := range encryptionAlgorithmDetails {
		if oid.Equal(details.oid) {
			return details.algo
		}
	}
	return UnknownEncryptionAlgorithm
}

var (
	oidEncryptionRSA          = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	oidEncryptionRC2_CBC      = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 2}
	oidEncryptionRC4          = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 4}
	oidEncryptionDES_EDE3_CBC = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 7}
	oidEncryptionRC5_CBC_Pad  = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 9}
	oidEncryptionID_DES_CDMF  = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 10}
)

var encryptionAlgorithmDetails = []struct {
	algo EncryptionAlgorithm
	oid  asn1.ObjectIdentifier
}{
	{RSA, oidEncryptionRSA},
	{RC2_CBC, oidEncryptionRC2_CBC},
	{RC4, oidEncryptionRC4},
	{DES_EDE3_CBC, oidEncryptionDES_EDE3_CBC},
	{RC5_CBC_Pad, oidEncryptionRC5_CBC_Pad},
	{ID_DES_CDMF, oidEncryptionID_DES_CDMF},
}
