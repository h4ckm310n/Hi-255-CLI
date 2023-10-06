package main

import (
	"Hi-255-CLI/hi255_grpc"
	"context"
	"fmt"
)

func ParseCommand(lineStr string) {
	line := []rune(lineStr)
	var args []string
	inQuote := false
	pos := 0
	n := len(line)
	curr := ""
	for pos <= n {
		if pos == n {
			args = append(args, curr)
		} else if line[pos] == '"' { // Quote
			if inQuote {
				inQuote = false
				args = append(args, curr)
				curr = ""
			} else {
				inQuote = true
			}
		} else if pos < n-1 && line[pos] == '\\' { // Special chars
			if line[pos+1] == ' ' || line[pos+1] == '"' || line[pos+1] == '\\' {
				curr += string(line[pos+1])
				pos += 1
			} else {
				curr += "\\"
			}
		} else if line[pos] == ' ' { // Space
			if inQuote {
				curr += " "
			} else {
				args = append(args, curr)
				curr = ""
			}
		} else {
			curr += string(line[pos])
		}

		pos += 1
	}

	switch args[0] {
	case "help":
		fmt.Println("Usage: \n" +
			"devices: List remote devices\n" +
			"file device_id filepath: Send a file\n" +
			"text device_id text_content: Send a text\n" +
			"greeting remote_ip: Send a greeting")
	case "devices":
		for _, device := range hi255_grpc.Devices {
			fmt.Println(device)
		}
	case "file":
		deviceID := args[1]
		filepath := args[2]
		_, err := hi255_grpc.GRPCClient.SendFile(context.Background(), &hi255_grpc.SendFileRequest{
			FilePath: filepath,
			RemoteId: deviceID,
		})
		fmt.Println(err)
	case "text":
		deviceID := args[1]
		text := args[2]
		_, err := hi255_grpc.GRPCClient.SendText(context.Background(), &hi255_grpc.SendTextRequest{
			Text:     text,
			RemoteId: deviceID,
		})
		fmt.Println(err)
	case "greeting":
		ip := args[1]
		_, err := hi255_grpc.GRPCClient.SendGreeting(context.Background(), &hi255_grpc.SendGreetingRequest{RemoteAddress: ip})
		fmt.Println(err)
	}
}
