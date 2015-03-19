package kangen

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func stoe(expire string) int64 {
	str := []byte(expire)
	format := regexp.MustCompile("([0-9]+)([smhd])")
	group := format.FindSubmatch(str)

	var result int64 = -1
	if len(group) == 3 {
		num, _ := strconv.ParseInt(string(group[1]), 10, 0)
		switch string(group[2]) {
		case "s":
			result = num
		case "m":
			result = num * int64(time.Minute) / int64(time.Second)
		case "h":
			result = num * int64(time.Hour) / int64(time.Second)
		case "d":
			result = num * int64(time.Hour) * 24 / int64(time.Second)
		}
	}

	return result
}

func etos(ttl int64) string {
	var result []byte

	if days := ttl / 86400; days > 0 {
		result = append(result, fmt.Sprintf("%dd", days)...)
		ttl %= 86400
	}
	if hours := ttl / 3600; hours > 0 {
		result = append(result, fmt.Sprintf("%dh", hours)...)
		ttl %= 3600
	}
	if minutes := ttl / 60; minutes > 0 {
		result = append(result, fmt.Sprintf("%dm", minutes)...)
		ttl %= 60
	}
	result = append(result, fmt.Sprintf("%ds", ttl)...)

	return string(result)
}
