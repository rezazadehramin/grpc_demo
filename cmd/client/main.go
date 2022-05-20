package main

import (
	"context"
	"demo_grpc/pkg/util"
	"demo_grpc/protos/everphone"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := everphone.NewEverphoneClient(conn)

	fmt.Println("Press ctrl+c to exit :D")

	for true {
		message := &everphone.EverphoneRandomTextInput{
			Text: util.RandStringRunes(10),
		}

		output, err := client.RandomText(context.Background(), message)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Output is: %s \n", output.Text)
		time.Sleep(3 * time.Second)
	}

}
