package main

import (
	"log"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

// CheckDB is const database name
const CheckDB = "health_check"

// Insert Result to InfluxdDB
func Insert(checkParam *HealthCheck, host *string) {
	// 	Initial Influxdb Client
	c, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: *host,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	//	Ceate batch point
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database:  CheckDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	//	tags
	tags := map[string]string{"target": checkParam.TARGET}
	fields := map[string]interface{}{"result": checkParam.STATUS}

	pt, err := influxdb.NewPoint("healthchecks", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
