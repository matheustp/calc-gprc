package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	calcpb "github.com/matheustp/calc-gprc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tls := true //Switch to run without certificates
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalcClient(cc)

	doUnary(c)
	doServerStream(c)
	doClientStream(c)
	doBidirectionalStream(c)
	doUnaryWithDeadline(c, 5*time.Second)
	doUnaryWithDeadline(c, 2*time.Second)
}

func doUnary(c calcpb.CalcClient) {
	req := calcpb.CalculateRequest{
		Num1:      10,
		Num2:      32,
		Operation: calcpb.Operation_SUM,
	}
	result, err := c.Calculate(context.Background(), &req)
	if err != nil {
		log.Panicf("Unary - Error when calling function: %v", err)
	}
	log.Println("Unary - Result:", result.GetResult())
}

func doServerStream(c calcpb.CalcClient) {
	req := &calcpb.PrimeNumberDecomposerRequest{
		Number: 42,
	}
	stream, err := c.PrimeNumberDecompose(context.Background(), req)
	if err != nil {
		log.Panicf("Server Stream - Error when calling function: %v", err)
	}
	result := []int32{}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("Server Stream - Error fetching message: %v", err)
		}
		result = append(result, msg.GetResult())
	}
	log.Printf("Server Stream - The prime number decomposers for %v are: %v", req.GetNumber(), result)
}

func doClientStream(c calcpb.CalcClient) {
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Client Stream - Error calling function: %v", err)
	}

	numToCalc := []*calcpb.ComputeAverageRequest{
		&calcpb.ComputeAverageRequest{
			Num: 2,
		},
		&calcpb.ComputeAverageRequest{
			Num: 10,
		},
		&calcpb.ComputeAverageRequest{
			Num: 2,
		},
		&calcpb.ComputeAverageRequest{
			Num: 3,
		},
	}

	for _, req := range numToCalc {
		log.Printf("Client Stream - Sending %v\n", req.GetNum())
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Client Stream - Error when receiving response: %v", err)
	}
	log.Printf("Client Stream - Result: %v", res.GetResult())
}

func doBidirectionalStream(c calcpb.CalcClient) {
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Panicf("Birectional Stream - Error calling function: %v", err)
	}
	wchan := make(chan struct{})
	go func() {
		numbers := []int32{1, 5, 3, 6, 2, 20}
		for _, n := range numbers {
			log.Printf("Birectional Stream - Sending %v\n", n)
			stream.Send(&calcpb.FindMaximumRequest{
				Num: n,
			})
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Birectional Stream - Error receiving message: %v", err)
				break
			}
			log.Printf("Birectional Stream - New maximum value: %v\n", res.GetMax())
		}
		close(wchan)
	}()
	<-wchan
}

func doUnaryWithDeadline(c calcpb.CalcClient, t time.Duration) {
	req := calcpb.CalculateWithDeadlineRequest{
		Num1:      10,
		Num2:      32,
		Operation: calcpb.Operation_MULTIPLY,
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	result, err := c.CalculateWithDeadline(ctx, &req)
	if err != nil {

		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Unary With Deadline - Timeout was hit! Deadline was exceeded")
			} else {
				log.Printf("Unary With Deadline - unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("Unary With Deadline - error while calling CalculateWithDeadline RPC: %v", err)
		}
		return
	}
	log.Println("Unary With Deadline - Result:", result.GetResult())
}
