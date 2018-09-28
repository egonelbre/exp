package smime

import "encoding/asn1"

// Attribute Object Identifiers
var (
	oidContentType      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 3}
	oidMessageDigest    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 4}
	oidSigningTime      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 5}
	oidCounterSignature = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 6}
)
