package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"

	pb "HandinTwo/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type hospital struct {
	pb.UnimplementedAggregatedShareServiceServer
	listOfReceivedShares []int64
}

func (h *hospital) SendAggregatedShare(ctx context.Context, msg *pb.AggregatedShare) (*pb.Acknowledge, error) {
	log.Printf("Received aggregated share: %v", msg.AggregatedShare)
	h.listOfReceivedShares = append(h.listOfReceivedShares, msg.AggregatedShare)

	if len(h.listOfReceivedShares) == 3 {
		log.Printf("Sum of shares: %d", h.sumShares())
	}

	return &pb.Acknowledge{Message: "Received aggregated share"}, nil
}

func (h *hospital) sumShares() int64 {
	var sum int64
	for _, share := range h.listOfReceivedShares {
		sum += share
	}
	return sum
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed the patient's certificate
	caCert, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// Create a certificate pool and append the CA certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load hospital server's certificate and private key
	hospitalServerCert, err := tls.LoadX509KeyPair("cert/hospital-server-cert.pem", "cert/hospital-server-key.pem")
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
	}

	// Create the credentials and return them
	config := &tls.Config{
		Certificates: []tls.Certificate{hospitalServerCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000") // Server listens on port 5000
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tlsCreds, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCreds),
	)
	pb.RegisterAggregatedShareServiceServer(grpcServer, &hospital{})

	fmt.Println("Hospital Server is running on port 5000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
