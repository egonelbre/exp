/*
Example how you would translate the ASN1 spec

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

func (d *asnSignerInfo) Unmarshal(tree *ber.Tree) error { return d.marshaler().Unmarshal(tree) }
func (d *asnSignerInfo) Marshal() (*ber.Tree, error) { return d.marshaler().Marshal() }
*/

package ber

// http://play.golang.org/p/MdL7k8-5ER

import (
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
)

// TODO: Optional/Check/Explicit should be all be a single type,
// otherwise it's hard to make them work nicely with nesting

type Marshaler interface {
	Marshal() (*Tree, error)
	Unmarshal(*Tree) error
}

type RawTree struct {
	Tree **Tree
}

func (m RawTree) Unmarshal(tree *Tree) error {
	*m.Tree = tree
	return nil
}

func (m RawTree) Marshal() (*Tree, error) {
	return *m.Tree, nil
}

type DER struct {
	V *[]byte
}

func (m DER) Unmarshal(tree *Tree) (err error) {
	*m.V, err = EncodeTree(tree)
	return
}

func (m DER) Marshal() (*Tree, error) {
	return DecodeTree(*m.V)
}

type ASN1RawValue struct {
	V *asn1.RawValue
}

func (m ASN1RawValue) Unmarshal(tree *Tree) (err error) {
	var data []byte
	if data, err = EncodeTree(tree); err != nil {
		return
	}

	_, err = asn1.Unmarshal(data, m.V)
	return
}

func (m ASN1RawValue) Marshal() (*Tree, error) {
	return ASN1Tree(m.V)
}

type Check struct {
	Class int
	Tag   int
	Sub   Marshaler
}

func (m Check) Unmarshal(tree *Tree) (err error) {
	if tree.Class != m.Class || tree.Tag != m.Tag {
		return fmt.Errorf("class/tag doesn't match: got %d/%d expected %d/%d", tree.Class, tree.Tag, m.Class, m.Tag)
	}
	return m.Sub.Unmarshal(tree)
}

func (m Check) Marshal() (*Tree, error) {
	tree, err := m.Sub.Marshal()
	if err != nil {
		return nil, err
	}

	tree.Class = m.Class
	tree.Tag = m.Tag
	return tree, nil
}

type Sequence []Marshaler

func (s Sequence) Unmarshal(tree *Tree) (err error) {
	children := tree.Children

	var si, ti int
	for si < len(s) {
		m := s[si]
		if opt, ok := m.(Optional); ok {
			if ti >= len(children) {
				si += 1
				continue
			}

			switch optsub := opt.Sub.(type) {
			case Check:
				sub := children[ti]
				if optsub.Class == sub.Class && optsub.Tag == sub.Tag {
					if err = optsub.Sub.Unmarshal(sub); err != nil {
						return
					}
					ti += 1
				}
				si += 1
			default:
				if err = optsub.Unmarshal(children[ti]); err != nil {
					return
				}
				ti += 1
				si += 1
			}
		} else {
			if ti >= len(children) {
				return fmt.Errorf("not enough children in tree")
			}

			sub := children[ti]
			if err = m.Unmarshal(sub); err != nil {
				return
			}

			ti += 1
			si += 1
		}
	}
	return
}

func (s Sequence) Marshal() (root *Tree, err error) {
	root = &Tree{
		Token: &Token{
			Kind:  Constructed,
			Class: Universal,
			Tag:   TagSequence,
		},
	}

	var sub *Tree
	for _, m := range s {
		if sub, err = m.Marshal(); err != nil {
			return
		}
		if _, optional := m.(Optional); optional && sub == nil {
			// we can safely skip this
		} else {
			if sub == nil {
				err = errors.New("marshaling returned nil")
				return
			}
			root.Children = append(root.Children, sub)
		}
	}
	return
}

type Choice []Marshaler

func (s Choice) Unmarshal(tree *Tree) (err error) {
	for _, m := range s {
		check, ok := m.(Check)
		if !ok {
			return fmt.Errorf("choice sequence doesn't contain check")
		}

		if check.Class == tree.Class && check.Tag == tree.Tag {
			err = check.Unmarshal(tree)
			return
		}
	}

	return fmt.Errorf("applicable choice element not found")
}

func (s Choice) Marshal() (*Tree, error) {
	// is there a way to handle this?
	panic("choice shoudln't be used while marshaling")
	return nil, nil
}

type Optional struct {
	Sub Marshaler
}

func (m Optional) Unmarshal(tree *Tree) (err error) {
	return m.Sub.Unmarshal(tree)
}

func isExplicit(s Marshaler) bool {
	_, ok := s.(Explicit)
	return ok
}

func isZero(tree *Tree) bool {
	return (tree == nil) ||
		(tree.Kind == Value && len(tree.Bytes) == 0) ||
		(tree.Kind == Constructed && len(tree.Children) == 0)
}

func (m Optional) Marshal() (tree *Tree, err error) {
	if tree, err = m.Sub.Marshal(); err != nil {
		return
	}

	if isZero(tree) {
		return nil, nil
	} else if isExplicit(m.Sub) && isZero(tree.Children[0]) {
		return nil, nil
	}
	return
}

type Explicit struct {
	Sub Marshaler
}

func (m Explicit) Unmarshal(tree *Tree) (err error) {
	if len(tree.Children) != 1 {
		return fmt.Errorf("not enough values for explicit type")
	}
	return m.Sub.Unmarshal(tree.Children[0])
}

func (m Explicit) Marshal() (tree *Tree, err error) {
	var sub *Tree
	if sub, err = m.Sub.Marshal(); err != nil {
		return
	}
	tree = &Tree{
		Token: &Token{
			Kind:  Constructed,
			Class: Context,
			Tag:   0,
		},
		Children: []*Tree{sub},
	}
	return
}

// TYPES

type OctetString struct {
	V *[]byte
}

func (m OctetString) Unmarshal(tree *Tree) error {
	val, err := tree.AsOctetString()
	*m.V = val
	return err
}

func (m OctetString) Marshal() (*Tree, error) {
	return &Tree{
		Token: &Token{
			Kind:  Value,
			Class: Universal,
			Tag:   TagOctetString,
			Bytes: *m.V,
		},
	}, nil
}

type Int64 struct {
	V *int64
}

func (m Int64) Unmarshal(tree *Tree) error {
	val, err := tree.AsInt64()
	*m.V = val
	return err
}

func (m Int64) Marshal() (*Tree, error) {
	return ASN1Tree(m.V)
}

type BigInt struct {
	V **big.Int
}

func (m BigInt) Unmarshal(tree *Tree) error {
	val, err := tree.AsBigInt()
	*m.V = val
	return err
}
func (m BigInt) Marshal() (*Tree, error) {
	return ASN1Tree(m.V)
}

type ObjectIdentifier struct {
	V *asn1.ObjectIdentifier
}

func (m ObjectIdentifier) Unmarshal(tree *Tree) error {
	val, err := tree.AsObjectIdentifier()
	*m.V = val
	return err
}
func (m ObjectIdentifier) Marshal() (*Tree, error) {
	return ASN1Tree(m.V)
}

func ASN1Tree(v interface{}) (*Tree, error) {
	bytes, err := asn1.Marshal(v)
	if err != nil {
		return nil, err
	}
	return DecodeTree(bytes)
}
