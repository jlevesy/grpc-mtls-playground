package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"experiments/mtls/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	clientCert  = "./dist/client.cert"
	clientKey   = "./dist/client.key"
	trustedCert = "./dist/ca.cert"
)

func main() {
	// Read client certificate and private key secret
	cert, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatal(err)
	}

	caCertPEM, err := ioutil.ReadFile(trustedCert)
	if err != nil {
		log.Fatal("Unable to load ca cert file: ", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		log.Fatal("Failed to append ca cert")
	}

	// Setup GRPC connection
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "localhost",
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial("localhost:4242", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// Ping the server
	rsp, err := api.NewSecurePingClient(conn).Echo(context.Background(), &api.Ping{})
	if err != nil {
		log.Fatalf("could not play ping pong: %v", err)
	}

	log.Printf("Sent a ping, received %+v", rsp)
}
