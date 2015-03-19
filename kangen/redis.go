package kangen

import (
	"github.com/garyburd/redigo/redis"
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

func setExpire(shorten string, expire string) {
	seconds := stoe(expire)
	if seconds == -1 {
		return
	}

	conn := connectRedis()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", "kangen:"+shorten, seconds)
	checkError(err)
}

func getExpire(shorten string) string {
	conn := connectRedis()
	defer conn.Close()

	reply, err := conn.Do("TTL", "kangen:"+shorten)
	ttl, err := redis.Int64(reply, err)
	checkError(err)

	if ttl == -1 {
		return ""
	}
	return etos(ttl)
}
