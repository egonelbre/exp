package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"path"
	"path/filepath"
	"time"
)

var (
	isClient = flag.Bool("client", false, "Is the client")
	addr     = flag.String("addr", "localhost:8080", "`address`")
)

func main() {
	flag.Parse()

	if filepath.Ext(flag.Arg(0)) == ".pem" {
		if *isClient {
			log.Fatal(NewCert(flag.Arg(0)))
		} else {
			log.Fatal(NewCert(flag.Arg(0)))
		}
	}

	if *isClient {
		log.Fatal(Client())
	} else {
		log.Fatal(Server())
	}
}

func echo(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%v", req.URL)
}

func Client() error {
	config := &tls.Config{
		RootCAs: x509.NewCertPool(),
	}

	// load the client certificates
	clientCert, err := tls.LoadX509KeyPair("client.pem", "client.pem")
	if err != nil {
		return err
	}
	config.Certificates = append(config.Certificates, clientCert)

	// load list of verified servers
	validServers, err := ioutil.ReadFile("server.pem")
	if err != nil {
		return err
	}
	config.RootCAs.AppendCertsFromPEM(validServers)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	resp, err := client.Get("https://" + path.Join(*addr, "hello", "world"))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("RESPONSE", string(body))
	return nil
}

func Server() error {
	config := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,

		ClientCAs: x509.NewCertPool(),
	}

	// load the server certificates
	serverCert, err := tls.LoadX509KeyPair("server.pem", "server.pem")
	if err != nil {
		return err
	}
	config.Certificates = append(config.Certificates, serverCert)

	// load the verified clients
	clientCAs, err := ioutil.ReadFile("client.pem")
	if err != nil {
		return err
	}
	config.ClientCAs.AppendCertsFromPEM(clientCAs)

	// create the server
	server := &http.Server{
		Addr:      *addr,
		Handler:   http.HandlerFunc(echo),
		TLSConfig: config,
	}

	// setup tcp server
	inner, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}

	// start serving
	listener := tls.NewListener(tcpKeepAliveListener{inner.(*net.TCPListener)}, config)
	return server.Serve(listener)
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func NewCert(pemfile string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("error generating new key: %s", err)
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(5 * 365 * 24 * time.Hour) // 5 years

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Hello"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature |
			x509.KeyUsageDataEncipherment |
			x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},

		IsCA: true,
		BasicConstraintsValid: true,

		DNSNames: []string{"localhost"},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %s", err)
	}

	var pemOut bytes.Buffer
	pem.Encode(&pemOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	pem.Encode(&pemOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return ioutil.WriteFile(pemfile, pemOut.Bytes(), 0644)
}
