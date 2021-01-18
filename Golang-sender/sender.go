package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michaelbironneau/asbclient"
)

func main() {
	log.Printf("Starting")

	client := asbclient.New(asbclient.Queue, os.Getenv("sb_namespace"), os.Getenv("sb_key_name"), os.Getenv("sb_key_value"))

	path := os.Getenv("sb_queue")
	var max_msg int
	fmt.Printf("Enter number of messages to send\n")
	fmt.Scanf("%d", &max_msg)
	for i := 0; i < max_msg; i++ {

		// log.Printf("Send: %d", i)
		err := client.Send(path, &asbclient.Message{
			Body: []byte(fmt.Sprintf("message %d", i)),
		})

		if err != nil {
			log.Printf("Send error: %s", err)
		} else {
			log.Printf("Sent: %d", i)
		}

		// time.Sleep(time.Millisecond * 5)
	}

}
