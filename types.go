package main

import "time"

type HealthCheck struct {
	URL    string `json:url`
	TARGET string `json:target`
	TIME   time.Time
	STATUS int
}
