package main

import (
	pb "gRPC-with-golang/proto"
	"io"
	"log"
)

func (s *helloServer) HelloBidirectionalStreaming(stream pb.GreetService_HelloBidirectionalStreamingServer) error {
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

}
