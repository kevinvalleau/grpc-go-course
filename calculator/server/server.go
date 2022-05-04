package main

import pb "github.com/kevinvalleau/grpc-go-course/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}