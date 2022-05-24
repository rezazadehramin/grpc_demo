package main

import (
	"context"
	"demo_grpc_client/pkg/util"
	"demo_grpc_client/protos/everphone"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("goserver:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := everphone.NewEverphoneClient(conn)

	fmt.Println("Press ctrl+c to exit :D")

	for true {
		ctx := context.Background()

		output, err := client.RandomText(ctx, &everphone.EverphoneRandomTextInput{
			Text: util.RandStringRunes(10),
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("Output is: %s \n", output.Text)
		time.Sleep(3 * time.Second)
	}

}
