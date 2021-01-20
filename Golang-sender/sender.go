package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michaelbironneau/asbclient"
)

func main() {
	log.Printf("Starting")
	var maxMsg int
	fmt.Printf("Enter number of messages to send\n")
	fmt.Scanf("%d", &maxMsg)
	client := asbclient.New(asbclient.Queue, os.Getenv("sb_namespace"), os.Getenv("sb_key_name"), os.Getenv("sb_key_value"))
	path := os.Getenv("sb_queue")
	for j := 1; j <= maxMsg; j++ {
		err := client.Send(path, &asbclient.Message{
			Body: []byte(fmt.Sprintf("message %d", j)),
		})
		if err != nil {
			log.Printf("Send error: %s", err)
		} else {
			log.Printf("Sent: %d", j)
		}
	}
}
