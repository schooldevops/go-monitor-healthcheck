package main

import (
	"log"
	"net/http"
	"time"
)

//	Check health of target
func Check(url string, target string, t time.Time, responseChannel chan<- HealthCheck) {
	log.Println("Current time: ", t)
	param := HealthCheck{url, target, t, 0000}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(param.URL)
	if err != nil {
		log.Println(err)
		param.STATUS = 500
	} else {
		defer resp.Body.Close()
		param.STATUS = resp.StatusCode
	}

	responseChannel <- param
}
