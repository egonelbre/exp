package smime

import (
	"bytes"
	"testing"
)

// Disables tests that use DianeDSSSignByCarlInherit
//   x509.ParseCertificate doesn't know how to handle it
// such parameterless are documented in RFC 3279 2.3.2
const skipinherit = false

// skip s-mime content tests for now
const skipmime = false

// plain RSA isn't implemented in x509
const skiprsa = false

/*
	Tests based on RFC 4134

2. Constants Used in the Examples

	- Alice is the creator of the message bodies in this document.
	- Bob is the recipient of the messages.
	- Carl is a CA.
	- Diane sometimes gets involved with these folks.
	- Erica also sometimes gets involved.

Sample Content:
	ExContent_bin

Alice's private keys:
	AlicePrivDSSSign_pri
	AlicePrivRSASign_pri

Bob's private key:
	BobPrivRSAEncrypt_pri

Carl's private keys:
	CarlPrivDSSSign_pri
	CarlPrivRSASign_pri

Diane's private keys:
	DianePrivDSSSign_pri
	DianePrivRSASignEncrypt_pri

Alice's certificates:
	AliceDSSSignByCarlNoInherit_cer
	AliceRSASignByCarl_cer

Bob's certificate:
	BobRSASignByCarl_cer

Carl's certificates:
	CarlDSSSelf_cer
	CarlRSASelf_cer

Diane's certificates:
	DianeDSSSignByCarlInherit_cer
	DianeRSASignByCarl_cer

Certificate Revocation Lists:
	CarlDSSCRLForAll_crl
	CarlDSSCRLForCarl_crl
	CarlDSSCRLEmpty_crl
	CarlRSACRLForAll_crl
	CarlRSACRLForCarl_crl
	CarlRSACRLEmpty_crl
*/

// 3.1 ContentInfo with Data Type, BER
func TestDataTypeBER(t *testing.T) {
	data, err := Parse(Section_3_1_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	content, ok := data.(Data)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(content))
	}
}

// 3.2 ContentInfo with Data Type, DER
func TestDataTypeDER(t *testing.T) {
	data, err := Parse(Section_3_2_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	content, ok := data.(Data)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(content))
	}
}

// 4.1 Basic Signed Content, DSS
//   A SignedData with no attribute certificates, signed by Alice using
//   DSS, just her certificate (not Carl's root cert), no CRL.  The
//   message is ExContent, and is included in the eContent.  There are no
//   signed or unsigned attributes.
func TestBasicSignedContentDSS(t *testing.T) {
	data, err := Parse(Section_4_1_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.2 Basic Signed Content, RSA
//   Same as 4.1, except using RSA signatures.  A SignedData with no
//   attribute certificates, signed by Alice using RSA, just her
//   certificate (not Carl's root cert), no CRL.  The message is
//   ExContent, and is included in the eContent.  There are no signed or
//   unsigned attributes.
func TestBasicSignedContentRSA(t *testing.T) {
	if skiprsa {
		return
	}

	data, err := Parse(Section_4_2_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.3 Basic Signed Content, Detached Content
//   Same as 4.1, except with no eContent.  A SignedData with no attribute
//   certificates, signed by Alice using DSS, just her certificate (not
//   Carl's root cert), no CRL.  The message is ExContent, but the
//   eContent is not included.  There are no signed or unsigned
//   attributes.
func TestBasicSignedContentDetachedContent(t *testing.T) {
	data, err := Parse(Section_4_3_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), []byte{}) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	// should we assign the content, or pass it as a param?
	signed.Content = ExContent_bin
	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.4 Fancier Signed Content
//   Same as 4.1, but includes Carl's root cert, Carl's CRL, some signed
//   and unsigned attributes (Countersignature by Diane).  A SignedData
//   with no attribute certificates, signed by Alice using DSS, her
//   certificate and Carl's root cert, Carl's DSS CRL.  The message is
//   ExContent, and is included in the eContent.  The signed attributes
//   are Content Type, Message Digest and Signing Time; the unsigned
//   attributes are content hint and counter signature.  The message
//   includes also Alice's RSA certificate.
func TestFancierSignedContent(t *testing.T) {
	data, err := Parse(Section_4_4_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.5. All RSA Signed Message
//   Same as 4.2, but includes Carl's RSA root cert (but no CRL).  A
//   SignedData with no attribute certificates, signed by Alice using RSA,
//   her certificate and Carl's root cert, no CRL.  The message is
//   ExContent, and is included in the eContent.  There are no signed or
//   unsigned attributes.
func TestAllRSASignedMessage(t *testing.T) {
	if skiprsa {
		return
	}
	data, err := Parse(Section_4_5_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.6. Multiple Signers
//   Similar to 4.1, but the message is also signed by Diane.  Two
//   signerInfos (one for Alice, one for Diane) with no attribute
//   certificates, each signed using DSS, Alice's and Diane's certificate
//   (not Carl's root cert), no CRL.  The message is ExContent, and is
//   included in the eContent.  There are no signed or unsigned
//   attributes.
func TestMultipleSigners(t *testing.T) {
	if skipinherit {
		return
	}

	data, err := Parse(Section_4_6_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.7. Signing Using SKI
//   Same as 4.1, but the signature uses the SKI instead of the
//   issuer/serial number in the cert.  A SignedData with no attribute
//   certificates, signed by Alice using DSS, just her certificate (not
//   Carl's root cert), identified by the SKI, no CRL.  The message is
//   ExContent, and is included in the eContent.  There are no signed or
//   unsigned attributes.
func TestSigningUsingSKI(t *testing.T) {
	data, err := Parse(Section_4_7_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.8.  S/MIME multipart/signed Message
//   A full S/MIME message, including MIME, that includes the body part
//   from 4.3 and the body containing the content of the message.
func TestSMIMEMultipartSignedMessage(t *testing.T) {
	if skipmime {
		return
	}
	// data Section_4_8_eml
	t.Fatalf("not implemented")
}

// 4.9.  S/MIME application/pkcs7-mime Signed Message
//   A full S/MIME message, including the MIME parts.
func TestSMIMEPKCS7SignedMessage(t *testing.T) {
	if skipmime {
		return
	}
	// data Section_4_9_eml
	t.Fatalf("not implemented")
}

// 4.10.  SignedData with Attributes
//   A SignedData message with the following list of signedAttributes:
//    -unknown OID
//    -contentHints
//    -smimeCapablilties
//    -securityLabel
//    -ContentReference
//    -smimeEncryptKeyPreference
//    -mlExpansionHistory
//    -EquivalentLabel
func TestSignedDataWithAttributes(t *testing.T) {
	data, err := Parse(Section_4_10_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), ExContent_bin) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err != nil {
		t.Fatalf("signature check failed: %v", err)
	}
}

// 4.11.  SignedData with Certificates Only
//   CA SignedData message with no content or signature, containing only
//   Alices's and Carl's certificates.
func TestSignedDataWithCertificates(t *testing.T) {
	data, err := Parse(Section_4_11_bin)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	signed, ok := data.(*SignedData)
	if !ok {
		t.Fatalf("content wrong type: %T", data)
	}

	if !bytes.Equal([]byte(signed.Content), []byte{}) {
		t.Fatalf("content is incorrect: \"%v\"", string(signed.Content))
	}

	if err := signed.CheckSignature(); err == nil {
		t.Fatalf("signature check must fail")
	}
}

/*
	Section_5_1_bin
	Section_5_2_bin
	Section_5_3_eml
	Section_6_0_bin
	Section_7_1_bin
	Section_7_2_bin
*/
