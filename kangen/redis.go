package kangen

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type kangenStruct struct {
	Shorten string `redis:"shorten"`
	URL     string `redis:"url"`
}

func connectRedis() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	checkError(err)

	return conn
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
