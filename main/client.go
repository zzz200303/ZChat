package main

import (
	"ZeZeIM/main/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"sync"
	"time"
)

func callPutStream(list []string, client hello.HelloServiceClient) {
	stream, err := client.PutStream(context.Background())
	if err != nil {
		fmt.Println("callPutStream err:", err)
	}
	for _, value := range list {
		if err := stream.Send(&hello.HelloRequest{Message: value}); err != nil {
			fmt.Println(err)
		}
		fmt.Println("client成功发送: " + value)
		time.Sleep(time.Second)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Message)
}

func callAllStream(list []string, client hello.HelloServiceClient) {
	stream, err := client.AllStream(context.Background())
	if err != nil {
		fmt.Println("callAllStream err:", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, value := range list {
			if err := stream.Send(&hello.HelloRequest{Message: value}); err != nil {
				fmt.Println(err)
			}
			fmt.Println("客户端发送了：" + value)
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(response.Message)
		}
	}()
	wg.Wait()
}

func callGetStream(message string, client hello.HelloServiceClient) {
	stream, err := client.GetStream(context.Background(), &hello.HelloRequest{Message: message})
	if err != nil {
		fmt.Println(err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("GetStream is done")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(response.Message)
	}
}

func main() {
	dial, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer dial.Close()
	dialClient := hello.NewHelloServiceClient(dial)
	callPutStream([]string{"客户端发送的第一条PutStream消息", "客户端发送的第二条PutStream消息"}, dialClient)
	callAllStream([]string{"客户端发送的第一条AllStream消息", "客户端发送的第二条AllStream消息"}, dialClient)
	callGetStream("客户端发送的GetStream消息", dialClient)
	echo, err := dialClient.Echo(context.Background(), &hello.HelloRequest{Message: "我好帅"})
	if err != nil {
		return
	}
	fmt.Println(echo.Message)
}
