package ber

type Token struct {
	Kind  Kind
	Class int
	Tag   int
	Bytes []byte
}

type Kind int

const (
	Value Kind = iota
	Constructed
	EndConstructed
)

const (
	Universal   = 0x00
	Application = 0x01
	Context     = 0x02
	Private     = 0x03
)

const (
	TagEOC              = 0x00
	TagBool             = 0x01
	TagInteger          = 0x02
	TagBitString        = 0x03
	TagOctetString      = 0x04
	TagNULL             = 0x05
	TagObjectIdentifier = 0x06
	TagObjectDescriptor = 0x07
	TagExternal         = 0x08
	TagRealFloat        = 0x09
	TagEnumerated       = 0x0a
	TagEmbeddedPDV      = 0x0b
	TagUTF8String       = 0x0c
	TagRelativeOID      = 0x0d
	TagSequence         = 0x10
	TagSet              = 0x11
	TagNumericString    = 0x12
	TagPrintableString  = 0x13
	TagT61String        = 0x14
	TagVideotexString   = 0x15
	TagIA5String        = 0x16
	TagUTCTime          = 0x17
	TagGeneralizedTime  = 0x18
	TagGraphicString    = 0x19
	TagVisibleString    = 0x1a
	TagGeneralString    = 0x1b
	TagUniversalString  = 0x1c
	TagCharacterString  = 0x1d
	TagBMPString        = 0x1e
)

// A StructuralError suggests that the ASN.1 data is valid, but the Go type
// which is receiving it doesn't match.
type StructuralError struct {
	Msg string
}

func (e StructuralError) Error() string { return "asn1: structure error: " + e.Msg }

// A SyntaxError suggests that the ASN.1 data is invalid.
type SyntaxError struct {
	Msg string
}

func (e SyntaxError) Error() string { return "asn1: syntax error: " + e.Msg }

func parseBase128Int(bytes []byte, initOffset int) (ret, offset int, err error) {
	offset = initOffset
	for shifted := 0; offset < len(bytes); shifted++ {
		if shifted > 4 {
			err = StructuralError{"base 128 integer too large"}
			return
		}
		ret <<= 7
		b := bytes[offset]
		ret |= int(b & 0x7f)
		offset++
		if b&0x80 == 0 {
			return
		}
	}
	err = SyntaxError{"truncated base 128 integer"}
	return
}

var unimplementedType = StructuralError{"unimplemented type"}
