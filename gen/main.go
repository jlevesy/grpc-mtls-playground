package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	log.Println("Generating the Root CA private key...")

	caPrivateKey, err := generateKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("generating the CA certificate...")

	caCertDER, err := generateCertificate(
		caPrivateKey.Public(),
		caPrivateKey,
		x509.Certificate{
			NotBefore:             time.Now(),
			NotAfter:              time.Now().Add(1200 * time.Hour),
			KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
			BasicConstraintsValid: true,
			IsCA:                  true,
		},
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Writing CA cert file...")

	if err = savePEM("dist/ca.cert", &pem.Block{Type: "CERTIFICATE", Bytes: caCertDER}); err != nil {
		log.Fatal(err)
	}

	log.Println("Parsing the root CA cert...")
	rootCACert, err := x509.ParseCertificate(caCertDER)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Generating the server private key...")
	serverPrivateKey, err := generateKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Generating the server certificate...")
	serverCertDER, err := generateCertificate(
		serverPrivateKey.Public(),
		caPrivateKey,
		x509.Certificate{
			NotBefore:   time.Now(),
			NotAfter:    time.Now().Add(1200 * time.Hour),
			KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
			ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

			BasicConstraintsValid: true,
			DNSNames:              []string{"test", "localhost"},
		},
		rootCACert,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err = savePEM(
		"dist/server.key",
		&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(serverPrivateKey)},
	); err != nil {
		log.Fatal(err)
	}

	if err = savePEM("dist/server.cert", &pem.Block{Type: "CERTIFICATE", Bytes: serverCertDER}); err != nil {
		log.Fatal(err)
	}

	log.Println("Generating the client private key...")
	clientPrivateKey, err := generateKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Generating the client certificate...")
	clientCertDER, err := generateCertificate(
		clientPrivateKey.Public(),
		caPrivateKey,
		x509.Certificate{
			NotBefore:             time.Now(),
			NotAfter:              time.Now().Add(1200 * time.Hour),
			KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			BasicConstraintsValid: true,
			DNSNames:              []string{"test", "localhost"},
		},
		rootCACert,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err = savePEM("dist/client.key", &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientPrivateKey)}); err != nil {
		log.Fatal(err)
	}

	if err = savePEM("dist/client.cert", &pem.Block{Type: "CERTIFICATE", Bytes: clientCertDER}); err != nil {
		log.Fatal(err)
	}

	log.Println("Certificates generated !")
}

func generateKeyPair() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 4096)
}

func generateCertificate(publicKey crypto.PublicKey, signerKey *rsa.PrivateKey, template x509.Certificate, parent *x509.Certificate) ([]byte, error) {
	var err error

	template.SerialNumber, err = rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %v", err)
	}

	if parent == nil {
		parent = &template
	}

	return x509.CreateCertificate(rand.Reader, &template, parent, publicKey, signerKey)
}

func savePEM(filename string, block *pem.Block) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, block)

}
