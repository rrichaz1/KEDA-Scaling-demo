package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/michaelbironneau/asbclient"
)
func checkmsg() {
	
	client := asbclient.New(asbclient.Queue, os.Getenv("sb_namespace"), os.Getenv("sb_key_name"), os.Getenv("sb_key_value"))

	path := os.Getenv("sb_queue")

	for {
		log.Printf("Peeking...")
		msg, err := client.PeekLockMessage(path, 30)

		if err != nil {
			log.Printf("Peek error: %s", err)
		} else {
			log.Printf("Peeked message: '%s'", string(msg.Body))
			err = client.DeleteMessage(msg)
			if err != nil {
				log.Printf("Delete error: %s", err)
			} else {
				log.Printf("Deleted message")
			}
		}
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second * time.Duration(rand.Intn(40)))
	}
}
func main() {
	log.Printf("Starting")
	go checkmsg()
	for {

	}

}
