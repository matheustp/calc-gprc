package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/reflection"

	calcpb "github.com/matheustp/calc-gprc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) ComputeAverage(stream calcpb.Calc_ComputeAverageServer) error {
	var total, i int32
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Result: float32(total) / float32(i),
			})
		}
		if err != nil {
			return err
		}
		total += res.GetNum()
		i++
	}
}

func (*server) FindMaximum(stream calcpb.Calc_FindMaximumServer) error {
	max := int32(0)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if res.GetNum() > max {
			max = res.GetNum()
			sendErr := stream.Send(&calcpb.FindMaximumResponse{
				Max: max,
			})
			if sendErr != nil {
				return sendErr
			}
		}
		log.Printf("Received: %v. Maximum is: %v\n", res.GetNum(), max)
	}
}

func (*server) PrimeNumberDecompose(req *calcpb.PrimeNumberDecomposerRequest, stream calcpb.Calc_PrimeNumberDecomposeServer) error {
	k := int32(2)
	n := req.GetNumber()
	for n > 1 {
		if n%k == 0 {
			res := &calcpb.PrimeNumberDecomposerResponse{
				Result: k,
			}
			stream.Send(res)
			n = n / k
		} else {
			k++
		}
	}
	return nil
}

func (*server) Calculate(ctx context.Context, req *calcpb.CalculateRequest) (*calcpb.CalculateResponse, error) {
	result := float32(0)
	switch req.GetOperation() {
	case calcpb.Operation_SUM:
		result = req.GetNum1() + req.GetNum2()
	case calcpb.Operation_DIFF:
		result = req.GetNum1() - req.GetNum2()
	case calcpb.Operation_MULTIPLY:
		result = req.GetNum1() * req.GetNum2()
	case calcpb.Operation_DIVIDE:
		if req.GetNum2() == 0 {
			return nil, status.Error(codes.InvalidArgument, "Cannot divide by ZERO")
		}
		result = req.GetNum1() / req.GetNum2()
	default:
		return nil, status.Error(codes.InvalidArgument, "You should set an operation")
	}
	return &calcpb.CalculateResponse{
		Result: result,
	}, nil
}

func (*server) CalculateWithDeadline(ctx context.Context, req *calcpb.CalculateWithDeadlineRequest) (*calcpb.CalculateWithDeadlineResponse, error) {
	//Wait for 3 seconds before start processing anything
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	result := float32(0)
	switch req.GetOperation() {
	case calcpb.Operation_SUM:
		result = req.GetNum1() + req.GetNum2()
	case calcpb.Operation_DIFF:
		result = req.GetNum1() - req.GetNum2()
	case calcpb.Operation_MULTIPLY:
		result = req.GetNum1() * req.GetNum2()
	case calcpb.Operation_DIVIDE:
		if req.GetNum2() == 0 {
			return nil, status.Error(codes.InvalidArgument, "Cannot divide by ZERO")
		}
		result = req.GetNum1() / req.GetNum2()
	default:
		return nil, status.Error(codes.InvalidArgument, "You should set an operation")
	}
	return &calcpb.CalculateWithDeadlineResponse{
		Result: result,
	}, nil
}

func main() {
	// Code to get file name and line when an error occurs
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error when starting to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	tls := true //Switch to false to use without certificate
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	calcpb.RegisterCalcServer(s, &server{})
	reflection.Register(s)

	go func() {
		log.Println("Starting server")
		if err := s.Serve(l); err != nil {
			log.Panicf("Error when starting to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Block until signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	l.Close()
	fmt.Println("End of Program")
}
