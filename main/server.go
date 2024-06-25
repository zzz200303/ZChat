package main

import (
	"ZeZeIM/main/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"sync"
)

type service struct {
	hello.UnimplementedHelloServiceServer
}

func (s *service) PutStream(stream hello.HelloService_PutStreamServer) error {
	for {
		rev, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("PutStream is done")
			err := stream.SendAndClose(&hello.HelloReply{Message: "PutStream is done"})
			if err != nil {
				return err
			}
			break //不加这个break会怎样
		}
		if err != nil {
			return err
		}
		fmt.Println("rev from client:", rev.Message)
	}
	return nil
}

func (s *service) GetStream(req *hello.HelloRequest, stream hello.HelloService_GetStreamServer) error {
	var stringList = []string{"This is the first message", "This is the second message"}
	for _, value := range stringList {
		if err := stream.Send(&hello.HelloReply{Message: value}); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) AllStream(stream hello.HelloService_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			rev, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Rev is done")
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("rev from client:", rev.Message)
		}
	}()
	go func() {
		defer wg.Done()
		list := []string{"第一条消息", "第二条消息", "第三条消息"}
		for _, value := range list {
			if err := stream.Send(&hello.HelloReply{Message: value}); err != nil {
				fmt.Println(err)
			}
		}
	}()
	wg.Wait()
	return nil
}

func (s *service) Echo(ctx context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{
		Message: "Echo: " + request.Message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &service{})
	fmt.Println("server listening at 5000")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
