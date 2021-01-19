package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michaelbironneau/asbclient"
)
var i int
func sendMsg(max_msg int) {
	client := asbclient.New(asbclient.Queue, os.Getenv("sb_namespace"), os.Getenv("sb_key_name"), os.Getenv("sb_key_value"))
	path := os.Getenv("sb_queue")
	for j := 0; j < max_msg; j++ {
		i++
		// log.Printf("Send: %d", i)
		err := client.Send(path, &asbclient.Message{
			Body: []byte(fmt.Sprintf("message %d", i)),
		})
		if err != nil {
			log.Printf("Send error: %s", err)
		} else {
			log.Printf("Sent: %d", i)
		}
	}
}
func main() {
	log.Printf("Starting")
	var max_msg int
	fmt.Printf("Enter number of messages to send per thread, note we use 10 threads\n")
	fmt.Scanf("%d", &max_msg)
	i = 0;
	for j:=0; j<9; j++ {
		go sendMsg(max_msg)
	}
	sendMsg(max_msg)
	for {
		if i >=100 {
			break
		}
	}
}
