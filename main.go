package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const defaultDb = "http://localhost:8086"

func main() {
	// dbHost := flag.String("-e", defaultDb, "HOST")
	dbHost := os.Getenv("DB_HOST")
	log.Println("env DB Host: ", dbHost)

	interval := os.Getenv("INTERVAL")
	log.Println("env Interval: ", interval)
	timeDuration, err := time.ParseDuration(interval)
	if err != nil {
		timeDuration = 2 * time.Second
	}

	ticker := time.NewTicker(timeDuration)
	defer ticker.Stop()

	targets := readJson()

	resultChannel := make(chan HealthCheck, len(targets))
	done := make(chan bool)

	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()
	for {
		select {
		case <-done:
			log.Println("Done!")
			return
		case t := <-ticker.C:
			for _, target := range targets {
				log.Println("Checking --- ", target.TARGET)
				go Check(target.URL, target.TARGET, t, resultChannel)
			}
			// Check("http://www.daum.net", "DAUM", t, resultChannel)

		case result := <-resultChannel:
			go Insert(&result, &dbHost)
		}

	}
}

func readJson() []HealthCheck {

	cwd, _ := os.Getwd()
	p := path.Join(cwd, "data", "targets.json")

	// Open our jsonFile
	jsonFile, err := os.Open(p)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened ./data/targets.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Println(err)
	}

	// we initialize our Users array
	var healthCheck []HealthCheck

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &healthCheck)

	return healthCheck
}
