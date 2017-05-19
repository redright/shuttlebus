package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func CreateNewClient() {
	fmt.Println("Preparing redis connection..")
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	fmt.Println("redis Connected")
	pubsub := client.PSubscribe("slocation-*")
	defer pubsub.Close()
	for {
		msgi, err := pubsub.Receive()
		if err != nil {
			fmt.Println("PubSub error:", err.Error())
			return
		}
		switch msg := msgi.(type) {
		case *redis.Message:
			fmt.Println("Received", msg.Payload, "on channel", msg.Channel)
		default:
			fmt.Println("Got control message", msg)
		}
	}
}
