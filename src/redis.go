package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func f(i, j int) (int, int) {
	return i, j
}

func f2(k, i, j int) (int, int) {
	return j + k, i + k
}

type Map struct {
	Key1 string "redis:key1"
	Key2 string "redis:key2"
}

func main() {
	c, err := redis.Dial("tcp", "192.168.1.101:6379")
	if err != nil {
		log.Fatal(err)
	}

	c.Send("MULTI")

	c.Send("HSET", "h1", "key1", "val2")
	c.Send("HSET", "h1", "key2", "val2")

	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")

	c.Send("HGETALL", "h1")

	// r, err := c.Do("EXEC")

	log.Println(redis.Values(c.Do("EXEC")))

	log.Println(redis.Strings(c.Do("HGETALL", "h1")))

	// ss, _ := redis.Strings(c.Do("HGETALL", "h1"))
	// var m Map
	// redis.ScanStruct(ss, m)


	// log.Println(m)

	// log.Println(f2(f(1, 2)))

}
