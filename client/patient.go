package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	pb "HandinTwo/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type patient struct {
	pb.UnimplementedShareSendingServiceServer
	patientName          string
	addressPort          string
	otherPatients        map[string]string // map of patient name to addressPort
	serverAddress        string
	listOfReceivedShares []int64
	privateInput         int64
}

func (p *patient) SendShare(ctx context.Context, msg *pb.Share) (*pb.Acknowledge, error) {
	log.Printf("Received share from other patient: %v", msg.ShareOfSecret)
	p.listOfReceivedShares = append(p.listOfReceivedShares, msg.ShareOfSecret)

	if len(p.listOfReceivedShares) == 3 {
		sum := p.sumShares()
		log.Printf("Sum of shares: %d", sum)
		p.SendAggregatedShareToHospital(context.Background(), sum)
	}

	return &pb.Acknowledge{Message: "Share received"}, nil
}

func (p *patient) sumShares() int64 {
	var sum int64
	for _, share := range p.listOfReceivedShares {
		sum += share
	}
	return sum
}

func (p *patient) StartPatientServer(wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", p.addressPort)
	if err != nil {
		log.Fatalf("[%s] Failed to listen on %s: %v", p.patientName, p.addressPort, err)
	}

	tlsCredentials, err := loadTLSCredentials(p.patientName)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	pb.RegisterShareSendingServiceServer(grpcServer, p)

	log.Printf("%s is running on %s", p.patientName, p.addressPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[%s] Failed to serve: %v", p.patientName, err)
	}
}

func (p *patient) SendShareToPatient(ctx context.Context, share int64, otherPatientName string) {
	tlsCreds, err := loadTLSCredentials(p.patientName)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(p.otherPatients[otherPatientName], grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		log.Fatalf("[%s] Failed to dial: %v", p.patientName, err)
	}
	defer conn.Close()

	log.Printf("Sending share to %s: %v", otherPatientName, share)

	client := pb.NewShareSendingServiceClient(conn)
	ack, err := client.SendShare(ctx, &pb.Share{ShareOfSecret: share})
	if err != nil {
		log.Fatalf("[%s] Failed to send share to %s: %v", p.patientName, otherPatientName, err)
	}

	log.Printf("Acknowledgment from %s: %s", otherPatientName, ack.Message)
}

func (p *patient) SendAggregatedShareToHospital(ctx context.Context, sumOfThreeShares int64) {
	tlsCreds, err := loadTLSCredentials(p.patientName)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(p.serverAddress, grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		log.Fatalf("[%s] Failed to dial hospital: %v", p.patientName, err)
	}
	defer conn.Close()

	log.Printf("Sending aggregated share to hospital: %v", sumOfThreeShares)

	client := pb.NewAggregatedShareServiceClient(conn)
	ack, err := client.SendAggregatedShare(ctx, &pb.AggregatedShare{AggregatedShare: sumOfThreeShares})
	if err != nil {
		log.Fatalf("[%s] Failed to send share to hospital: %v", p.patientName, err)
	}

	log.Printf("Acknowledgment from Hospital: %s", ack.Message)
}

func (p *patient) calculateShares() (int64, int64, int64) {
	// Generate two random ints
	share1 := rand.Int63()
	share2 := rand.Int63()

	// Calculate the third share
	share3 := p.privateInput - share1 - share2

	return share1, share2, share3
}

func loadTLSCredentials(patientName string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed hospital's certificate
	caCert, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// Create a certificate pool and append the CA certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	certFile := filepath.Join("cert", fmt.Sprintf("%s-cert.pem", patientName))
	keyFile := filepath.Join("cert", fmt.Sprintf("%s-key.pem", patientName))

	// Load patient's certificate and private key
	patientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return them
	config := &tls.Config{
		Certificates: []tls.Certificate{patientCert},
		RootCAs:      certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	patientName := flag.String("name", "", "Name of the patient")
	addressPort := flag.String("address", "", "Addressport of the patient")
	privateInput := flag.Int64("input", 0, "Private input of the patient")
	flag.Parse()

	otherPatients := map[string]string{
		"Alice":   "localhost:5001",
		"Bob":     "localhost:5002",
		"Charlie": "localhost:5003",
	}

	delete(otherPatients, *patientName)

	patient := &patient{
		patientName:          *patientName,
		addressPort:          *addressPort,
		otherPatients:        otherPatients,
		serverAddress:        "localhost:5000",
		listOfReceivedShares: []int64{},
		privateInput:         *privateInput,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go patient.StartPatientServer(&wg)

	time.Sleep(10 * time.Second)

	share1, share2, share3 := patient.calculateShares()
	log.Printf("Calculated shares: %v, %v, %v", share1, share2, share3)

	if patient.patientName == "Alice" {
		patient.listOfReceivedShares = append(patient.listOfReceivedShares, share1)
		time.Sleep(10 * time.Second)
		patient.SendShareToPatient(context.Background(), share2, "Bob")
		patient.SendShareToPatient(context.Background(), share3, "Charlie")
	} else if patient.patientName == "Bob" {
		patient.listOfReceivedShares = append(patient.listOfReceivedShares, share2)
		time.Sleep(10 * time.Second)
		patient.SendShareToPatient(context.Background(), share1, "Alice")
		patient.SendShareToPatient(context.Background(), share3, "Charlie")
	} else if patient.patientName == "Charlie" {
		patient.listOfReceivedShares = append(patient.listOfReceivedShares, share3)
		patient.SendShareToPatient(context.Background(), share1, "Alice")
		patient.SendShareToPatient(context.Background(), share2, "Bob")
	}

	// Keep the server running
	wg.Wait()
}
