package main

import (
	"net/http"
	"time"
)

func timeForRequest(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(url1 string, url2 string) string {
	time1 := timeForRequest(url1)
	time2 := timeForRequest(url2)
	if time1 > time2 {
		return url2
	} else {
		return url1
	}
}
