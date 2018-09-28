package ber

import (
	"encoding/asn1"
	"fmt"
	"math/big"
	"time"
)

func (v *Token) mustBeValue() {
	if v.Kind != Value {
		panic(fmt.Sprintf("invalid call on non-kind: %#v", v))
	}
}

func (v *Token) Universal() (ret interface{}, err error) {
	v.mustBeValue()
	if v.Class != Universal {
		err = SyntaxError{"not an universal value"}
		return
	}

	switch v.Tag {
	case TagBool:
		ret, err = v.AsBool()
	case TagInteger:
		ret, err = v.AsBigInt()
	case TagBitString:
		ret, err = v.AsBitString()
	case TagOctetString:
		ret, err = v.AsOctetString()
	case TagNULL:
		ret = struct{}{}
	case TagObjectIdentifier:
		ret, err = v.AsObjectIdentifier()
	// case TagObjectDescriptor:
	// case TagExternal:
	// case TagRealFloat:
	case TagEnumerated:
		ret, err = v.AsInt64()
	// case TagEmbeddedPDV:
	case TagUTF8String:
		ret, err = v.AsUTF8String()
	// case TagRelativeOID:
	// case TagSequence:
	// case TagSet:
	case TagNumericString:
		ret, err = v.AsNumericString()
	case TagPrintableString:
		ret, err = v.AsPrintableString()
	case TagT61String:
		ret, err = v.AsT61String()
	// case TagVideotexString:
	case TagIA5String:
		ret, err = v.AsIA5String()
	case TagUTCTime:
		ret, err = v.AsUTCTime()
	case TagGeneralizedTime:
		ret, err = v.AsGeneralizedTime()
	//case TagGraphicString:
	case TagVisibleString:
		ret, err = v.AsVisibleString()
	case TagGeneralString:
		ret, err = v.AsUTF8String()
	case TagUniversalString:
		ret, err = v.AsUTF8String()
	case TagCharacterString:
		ret, err = v.AsUTF8String()
	//case TagBMPString:
	default:
		err = unimplementedType
	}

	return
}

// BOOLEAN

func (v *Token) AsBool() (ret bool, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	if len(bytes) != 1 {
		err = SyntaxError{"invalid boolean"}
		return
	}

	// DER demands that "If the encoding represents the boolean value TRUE,
	// its single contents octet shall have all eight bits set to one."
	// Thus only 0 and 255 are valid encoded values.
	switch bytes[0] {
	case 0:
		ret = false
	case 1:
		ret = true
	default:
		err = SyntaxError{"invalid boolean"}
	}
	return
}

// INTEGER

// AsInt64 treats the given bytes as a big-endian, signed integer and
// returns the result.
func (v *Token) AsInt64() (ret int64, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	if len(bytes) > 8 {
		// We'll overflow an int64 in this case.
		err = StructuralError{"integer too large"}
		return
	}
	for bytesRead := 0; bytesRead < len(bytes); bytesRead++ {
		ret <<= 8
		ret |= int64(bytes[bytesRead])
	}

	// Shift up and down in order to sign extend the result.
	ret <<= 64 - uint8(len(bytes))*8
	ret >>= 64 - uint8(len(bytes))*8
	return
}

// AsInt32 treats the given bytes as a big-endian, signed integer and returns
// the result.
func (v *Token) AsInt32() (ret int32, err error) {
	v.mustBeValue()
	ret64, err := v.AsInt64()
	if err != nil {
		return 0, err
	}
	if ret64 != int64(int32(ret64)) {
		return 0, StructuralError{"integer too large"}
	}
	return int32(ret64), nil
}

var bigOne = big.NewInt(1)

// parseBigInt treats the given bytes as a big-endian, signed integer and returns
// the result.
func (v *Token) AsBigInt() (*big.Int, error) {
	v.mustBeValue()
	bytes := v.Bytes
	ret := new(big.Int)
	if len(bytes) > 0 && bytes[0]&0x80 == 0x80 {
		// This is a negative number.
		notBytes := make([]byte, len(bytes))
		for i := range notBytes {
			notBytes[i] = ^bytes[i]
		}
		ret.SetBytes(notBytes)
		ret.Add(ret, bigOne)
		ret.Neg(ret)
		return ret, nil
	}
	ret.SetBytes(bytes)
	return ret, nil
}

// BIT STRING

// AsBitString parses an ASN.1 bit string from the given byte slice and returns it.
func (v *Token) AsBitString() (ret asn1.BitString, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	if len(bytes) == 0 {
		err = SyntaxError{"zero length BIT STRING"}
		return
	}
	paddingBits := int(bytes[0])
	if paddingBits > 7 ||
		len(bytes) == 1 && paddingBits > 0 {
		err = SyntaxError{"invalid padding bits in BIT STRING"}
		return
	}
	ret.BitLength = (len(bytes)-1)*8 - paddingBits
	ret.Bytes = bytes[1:]

	// remove the padding bits
	last := len(ret.Bytes) - 1
	ret.Bytes[last] = (ret.Bytes[last] >> uint(paddingBits)) << uint(paddingBits)
	return
}

func (v *Token) AsOctetString() (ret []byte, err error) {
	v.mustBeValue()
	return v.Bytes, nil
}

// AsObjectIdentifier parses an OBJECT IDENTIFIER from the given bytes and
// returns it. An object identifier is a sequence of variable length integers
// that are assigned in a hierarchy.
func (val *Token) AsObjectIdentifier() (s asn1.ObjectIdentifier, err error) {
	val.mustBeValue()
	bytes := val.Bytes
	if len(bytes) == 0 {
		err = SyntaxError{"zero length OBJECT IDENTIFIER"}
		return
	}

	// In the worst case, we get two elements from the first byte (which is
	// encoded differently) and then every varint is a single byte long.
	s = make(asn1.ObjectIdentifier, len(bytes)+1)

	// The first varint is 40*Token1 + value2:
	// According to this packing, value1 can take the values 0, 1 and 2 only.
	// When value1 = 0 or value1 = 1, then value2 is <= 39. When value1 = 2,
	// then there are no restrictions on value2.
	v, offset, err := parseBase128Int(bytes, 0)
	if err != nil {
		return
	}
	if v < 80 {
		s[0] = v / 40
		s[1] = v % 40
	} else {
		s[0] = 2
		s[1] = v - 80
	}

	i := 2
	for ; offset < len(bytes); i++ {
		v, offset, err = parseBase128Int(bytes, offset)
		if err != nil {
			return
		}
		s[i] = v
	}
	s = s[0:i]
	return
}

// UTCTime

func (v *Token) AsUTCTime() (ret time.Time, err error) {
	v.mustBeValue()
	s := string(v.Bytes)
	ret, err = time.Parse("0601021504Z0700", s)
	if err != nil {
		ret, err = time.Parse("060102150405Z0700", s)
	}
	if err == nil && ret.Year() >= 2050 {
		// UTCTime only encodes times prior to 2050. See https://tools.ietf.org/html/rfc5280#section-4.1.2.5.1
		ret = ret.AddDate(-100, 0, 0)
	}

	return
}

// AsGeneralizedTime parses the GeneralizedTime from the given byte slice
// and returns the resulting time.
func (v *Token) AsGeneralizedTime() (ret time.Time, err error) {
	v.mustBeValue()
	return time.Parse("20060102150405Z0700", string(v.Bytes))
}

// NumericString
func (v *Token) AsNumericString() (ret string, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	for _, b := range bytes {
		if !isNumeric(b) {
			err = SyntaxError{"NumericString contains invalid character"}
			return
		}
	}
	ret = string(bytes)
	return
}

// isNumeric returns true iff the given b is in the ASN.1 NumericString set.
func isNumeric(b byte) bool {
	return '0' <= b && b <= '9' || b == ' '
}

// AsPrintableString parses a ASN.1 PrintableString from the given byte
// array and returns it.
func (v *Token) AsPrintableString() (ret string, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	for _, b := range bytes {
		if !isPrintable(b) {
			err = SyntaxError{"PrintableString contains invalid character"}
			return
		}
	}
	ret = string(bytes)
	return
}

// isPrintable returns true iff the given b is in the ASN.1 PrintableString set.
func isPrintable(b byte) bool {
	return 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z' ||
		'0' <= b && b <= '9' ||
		'\'' <= b && b <= ')' ||
		'+' <= b && b <= '/' ||
		b == ' ' ||
		b == ':' ||
		b == '=' ||
		b == '?' ||
		// This is technically not allowed in a PrintableString.
		// However, x509 certificates with wildcard strings don't
		// always use the correct string type so we permit it.
		b == '*'
}

// VisibleString

// AsVisibleString parses a ASN.1 VisibleString from the given byte
// array and returns it.
func (v *Token) AsVisibleString() (ret string, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	for _, b := range bytes {
		if b < 0x20 {
			err = SyntaxError{"VisibleString contains invalid character"}
			return
		}
	}
	ret = string(bytes)
	return
}

// IA5String

// parseIA5String parses a ASN.1 IA5String (ASCII string) from the given
// byte slice and returns it.
func (v *Token) AsIA5String() (ret string, err error) {
	v.mustBeValue()
	bytes := v.Bytes
	for _, b := range bytes {
		if b >= 0x80 {
			err = SyntaxError{"IA5String contains invalid character"}
			return
		}
	}
	ret = string(bytes)
	return
}

// T61String

// parseT61String parses a ASN.1 T61String (8-bit clean string) from the given
// byte slice and returns it.
func (v *Token) AsT61String() (ret string, err error) {
	v.mustBeValue()
	return string(v.Bytes), nil
}

// UTF8String

// parseUTF8String parses a ASN.1 UTF8String (raw UTF-8) from the given byte
// array and returns it.
func (v *Token) AsUTF8String() (ret string, err error) {
	v.mustBeValue()
	return string(v.Bytes), nil
}
