package kangen

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Exist(shorten string) bool {
	conn := connectRedis()
	defer conn.Close()

	reply, err := conn.Do("EXISTS", "kangen:"+shorten)
	status, err := redis.Int(reply, err)

	return status == 1
}

func Get(shorten string) string {
	conn := connectRedis()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("HGETALL", "kangen:"+shorten))
	if err != nil {
		return ""
	}

	var k kangenStruct
	redis.ScanStruct(reply, &k)

	return k.URL
}

func Add(shorten string, url string, expire string) string {
	conn := connectRedis()
	defer conn.Close()

	if Exist(shorten) {
		return fmt.Sprintf("[EXIST] %s -> %s\n", shorten, Get(shorten))
	}

	var k kangenStruct = kangenStruct{
		Shorten: shorten,
		URL:     url,
	}

	_, err := conn.Do("HMSET", redis.Args{}.Add("kangen:"+shorten).AddFlat(&k)...)
	checkError(err)
	setExpire(shorten, expire)

	return fmt.Sprintf("[ADD] %s -> %s\n", shorten, url)
}

func Remove(shorten string) string {
	conn := connectRedis()
	defer conn.Close()

	if !Exist(shorten) {
		return fmt.Sprintf("[NOT FOUND] %s\n", shorten)
	}
	url := Get(shorten)

	_, err := conn.Do("DEL", "kangen:"+shorten)
	checkError(err)

	return fmt.Sprintf("[REMOVE] %s -> %s\n", shorten, url)
}

func List() string {
	conn := connectRedis()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("KEYS", "kangen:*"))
	checkError(err)

	var keys []string
	redis.ScanSlice(reply, &keys)
	if len(keys) == 0 {
		return "[NOTHING]\n"
	}

	var result []byte
	var k kangenStruct
	for i := range keys {
		reply, err := redis.Values(conn.Do("HGETALL", keys[i]))
		checkError(err)

		redis.ScanStruct(reply, &k)
		if expire := getExpire(k.Shorten); expire == "" {
			result = append(result, fmt.Sprintf("%s -> %s\n", k.Shorten, k.URL)...)
		} else {
			result = append(result, fmt.Sprintf("%s -> %s (%s)\n", k.Shorten, k.URL, expire)...)
		}
	}

	return string(result)
}
