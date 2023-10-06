package main

import (
	"Hi-255-CLI/hi255_grpc"
	"bufio"
	"os"
)

func main() {
	hi255_grpc.InitGRPCClient()
	go hi255_grpc.FetchDevices()
	go hi255_grpc.FetchMessages()
	scanner := bufio.NewScanner(os.Stdin)
	cmd := ""
	for {
		if scanner.Scan() {
			cmd = scanner.Text()
			if cmd == "exit" {
				break
			}
			ParseCommand(cmd)
		}
	}
}
