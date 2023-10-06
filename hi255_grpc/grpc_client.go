package hi255_grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

var GRPCClient ServiceClient

type Device struct {
	ID       string
	Name     string
	Addr     string
	Platform string
}

var Devices map[string]*Device

func InitGRPCClient() {
	homeDir, _ := os.UserHomeDir()
	conn, err := grpc.Dial("unix://"+homeDir+"/.hi255/"+"hi255.sock", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	GRPCClient = NewServiceClient(conn)
}

func FetchDevices() {
	stream, err := GRPCClient.FetchRemoteDevices(context.Background(), &Empty{})
	Devices = make(map[string]*Device)
	if err != nil {
		return
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		devices := resp.RemoteDevices
		for _, device := range devices {
			Devices[device.Id] = &Device{
				ID:       device.Id,
				Name:     device.Name,
				Addr:     device.Address,
				Platform: device.Platform,
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func FetchMessages() {
	stream, err := GRPCClient.FetchMessages(context.Background(), &Empty{})
	if err != nil {
		return
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		messages := resp.Messages
		for _, m := range messages {
			fmt.Println(m)
		}
		time.Sleep(3 * time.Second)
	}
}
