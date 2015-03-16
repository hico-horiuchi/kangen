package kangen

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func connectRedis() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func getURL(conn redis.Conn, shorten string) string {
	reply, err := conn.Do("HGET", "kangen", shorten)
	url, err := redis.String(reply, err)

	if err != nil {
		return ""
	}

	return url
}

func Add(shorten string, url string) string {
	conn := connectRedis()
	reply, err := conn.Do("HSET", "kangen", shorten, url)
	status, err := redis.Int(reply, err)

	if err != nil {
		log.Fatalln(err)
	}

	var result string
	if status == 0 {
		result = fmt.Sprintf("[EXIST] %s -> %s\n", shorten, getURL(conn, shorten))
	} else {
		result = fmt.Sprintf("[ADD] %s -> %s\n", shorten, url)
	}

	defer conn.Close()
	return result
}

func Remove(shorten string) string {
	conn := connectRedis()
	reply, err := conn.Do("HDEL", "kangen", shorten)
	status, err := redis.Int(reply, err)

	if err != nil {
		log.Fatalln(err)
	}

	var result string
	if status == 0 {
		result = fmt.Sprintf("[NOT FOUND] %s\n", shorten)
	} else {
		result = fmt.Sprintf("[REMOVE] %s\n", shorten)
	}

	defer conn.Close()
	return result
}

func List() string {
	conn := connectRedis()
	reply, err := conn.Do("HGETALL", "kangen")
	arr, err := redis.StringMap(reply, err)

	if err != nil {
		log.Fatalln(err)
	} else if len(arr) == 0 {
		return "[NOTHING]\n"
	}

	var result []byte
	for k, v := range arr {
		result = append(result, fmt.Sprintf("%s -> %s\n", k, v)...)
	}

	defer conn.Close()
	return string(result)
}
