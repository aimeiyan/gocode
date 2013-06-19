package main

import (
	"github.com/hoisie/redis"
	"log"
)

type AA struct {
	Key3, Key4 string
	Key2       string "key1"
	Key1       string "key2"
}

func main() {
	client := redis.Client{Addr: "192.168.1.101:6379"}
	// client.Addr = "127.0.0.1:6379"

	client.Hmset("h1", map[string]string{"Key3": "abc", "Key4": "sdfs",
		"key1": "abc", "key2": "sdfs"})

	m := make(map[string]string)
	client.Hgetall("h1", m)
	log.Println(m)

	a := AA{}
	client.Hgetall("h1", &a)
	log.Println(a)
}
