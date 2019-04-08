package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net"

	"experiments/mtls/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var insecure = flag.Bool("insecure", false, "start the server in insecure mode")

const (
	serverCert  = "./dist/server.cert"
	serverKey   = "./dist/server.key"
	trustedCert = "./dist/ca.cert"
)

type apiHandler func(context.Context, *api.Ping) (*api.Pong, error)

func (a apiHandler) Echo(ctx context.Context, ping *api.Ping) (*api.Pong, error) {
	return a(ctx, ping)
}

func main() {
	log.Println("Starting...")
	flag.Parse()

	// Open a channel on the address for listening
	listener, err := net.Listen("tcp", ":4242")
	if err != nil {
		log.Fatalf("Could not listen on 4242: %v\n", err)
	}

	log.Println("Listening on 4242")

	// Create a new GRPC server with the credentials
	server := grpc.NewServer(grpcOptions()...)

	api.RegisterSecurePingServer(server, apiHandler(func(ctx context.Context, ping *api.Ping) (*api.Pong, error) {
		log.Println("Received an PING")
		return &api.Pong{}, nil
	}))

	log.Println("Created GRPC server, going to serve now")

	// Listen
	if err := server.Serve(listener); err != nil {
		log.Fatalf("grpc serve error: %v\n", err)
	}
}

func grpcOptions() []grpc.ServerOption {
	var options []grpc.ServerOption
	// If insecure mode is on, do not enable TLS.
	if *insecure {
		return options
	}

	serverCert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		log.Fatal("Cannot be loaded the certificate: ", err)
	}
	log.Println("Server cert loaded")

	caCertPEM, err := ioutil.ReadFile(trustedCert)
	if err != nil {
		log.Fatal("Unable to load ca cert file: ", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		log.Fatal("Failed to append CA cert")
	}

	log.Println("Client cert added to cert pool")

	// Create the TLS configuration to pass to the GRPC server
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
	})

	options = append(options, grpc.Creds(creds))

	return options
}
