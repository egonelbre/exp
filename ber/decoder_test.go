package ber

import (
	"bytes"
	"encoding/asn1"
	"io"
	"reflect"
	"testing"
)

type testcase struct {
	in     []byte
	expect []expectation
}
type testcases []testcase

type expectation struct {
	Kind  Kind
	Class int
	Tag   int
	Value interface{}
}

func TestBitString(t *testing.T) {
	runTestCases(t, testcases{
		testcase{ // simple
			in: []byte{0x03, 0x04, 0x06, 0x6e, 0x5d, 0xc0},
			expect: []expectation{
				{Value, Universal, TagBitString, asn1.BitString{[]byte{0x6E, 0x5D, 0xC0}, 18}},
			},
		},
		testcase{ // "100000" padding
			in: []byte{0x03, 0x04, 0x06, 0x6e, 0x5d, 0xe0},
			expect: []expectation{
				{Value, Universal, TagBitString, asn1.BitString{[]byte{0x6E, 0x5D, 0xC0}, 18}},
			},
		},
		testcase{ // long form length octets
			in: []byte{0x03, 0x81, 0x04, 0x06, 0x6e, 0x5d, 0xe0},
			expect: []expectation{
				{Value, Universal, TagBitString, asn1.BitString{[]byte{0x6E, 0x5D, 0xC0}, 18}},
			},
		},
		testcase{ // constructed
			in: []byte{0x23, 0x09,
				0x03, 0x03, 0x00, 0x6e, 0x5d,
				0x03, 0x02, 0x06, 0xc0},
			expect: []expectation{
				{Constructed, Universal, TagBitString, nil},
				{Value, Universal, TagBitString, asn1.BitString{[]byte{0x6E, 0x5D}, 16}},
				{Value, Universal, TagBitString, asn1.BitString{[]byte{0xC0}, 2}},
				{EndConstructed, Universal, TagBitString, nil},
			},
		},
	})
}

func TestCMS(t *testing.T) {
	runTestCases(t, testcases{
		testcase{ // simple
			in: []byte{0x30, 0x80, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x07, 0x01, 0xa0, 0x80, 0x24, 0x80, 0x04, 0x04, 0x54,
				0x68, 0x69, 0x73, 0x04, 0x18, 0x20, 0x69, 0x73, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expect: []expectation{
				{Constructed, Universal, TagSequence, nil},
				{Value, Universal, TagObjectIdentifier, asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}},
				{Constructed, Context, 0, nil},
				{Constructed, Universal, TagOctetString, nil},
				{Value, Universal, TagOctetString, []byte("This")},
				{Value, Universal, TagOctetString, []byte(" is some sample content.")},
				{EndConstructed, Universal, TagOctetString, nil},
				{EndConstructed, Context, 0, nil},
				{EndConstructed, Universal, TagSequence, nil},
			},
		},
		testcase{ // "100000" padding
			in: []byte{0x30, 0x2b, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x07, 0x01, 0xa0, 0x1e, 0x04, 0x1c, 0x54, 0x68, 0x69,
				0x73, 0x20, 0x69, 0x73, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x63, 0x6f, 0x6e,
				0x74, 0x65, 0x6e, 0x74, 0x2e},
			expect: []expectation{
				{Constructed, Universal, TagSequence, nil},
				{Value, Universal, TagObjectIdentifier, asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}},
				{Constructed, Context, 0, nil},
				{Value, Universal, TagOctetString, []byte("This is some sample content.")},
				{EndConstructed, Context, 0, nil},
				{EndConstructed, Universal, TagSequence, nil},
			},
		},
	})
}

func TestEmptySet(t *testing.T) {
	runTestCases(t, testcases{
		testcase{
			in: []byte{0x31, 0x00},
			expect: []expectation{
				{Constructed, Universal, TagSet, nil},
				{EndConstructed, Universal, TagSet, nil},
			},
		},
	})
}

func runTestCase(i int, test testcase, t *testing.T) {
	d := NewDecoder(bytes.NewReader(test.in))
	for j, expect := range test.expect {
		token, err := d.Token()
		if err != nil {
			t.Fatalf("error while parsing: %v", err)
		}

		if token.Kind != expect.Kind ||
			token.Class != expect.Class ||
			token.Tag != expect.Tag {
			t.Fatalf("wrong class/kind/tag at %d %d: got %#v expected %#v", i, j, token, expect)
		}

		if token.Kind == Value {
			value, err := token.Universal()
			if err != nil {
				t.Fatalf("error at %d %d had %#v: %v", i, j, value, err)
			}

			if !reflect.DeepEqual(value, expect.Value) {
				t.Fatalf("error at %d %d: got %#v expected %#v", i, j, value, expect.Value)
			}
		}
	}

	token, err := d.Token()
	if err != io.EOF {
		t.Fatalf("didn't get EOF, got %#v %#v", token, err)
	}
}

func runTestCases(t *testing.T, cases testcases) {
	for i, test := range cases {
		runTestCase(i, test, t)
	}
}
