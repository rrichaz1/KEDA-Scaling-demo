package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Azure/azure-service-bus-go"
)

func main() {
	log.Printf("Starting")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := os.Getenv("queueConnString")
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable SERVICEBUS_CONNECTION_STRING not set")
		return
	}

	// Create a client to communicate with a Service Bus Namespace.
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a client to communicate with the queue. (The queue must have already been created, see `QueueManager`)
	q, err := ns.NewQueue(os.Getenv("sb_queue"))
	if err != nil {
		fmt.Println("FATAL: ", err)
		return
	}

	for {
		err = q.ReceiveOne(
			ctx,
			servicebus.HandlerFunc(func(ctx context.Context, message *servicebus.Message) error {
				fmt.Println(string(message.Data))
				return message.Complete(ctx)
			}))
		if err != nil {
			fmt.Println("FATAL: ", err)
			return
		}

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second * time.Duration(rand.Intn(40)))
	}

}
