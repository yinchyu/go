package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("http://42.193.190.143:8086", "zhCU6HBUbIGx2fi4PBByV0cDW9KBzsrlwSHPNRojKOZmTVYKBS7iGSSyCfyWunqkQ8PNXqZpEa2XK9JwPaXtMA==")
	// always close client at the end
	defer client.Close()
	// get non-blocking write client
	writeAPI := client.WriteAPI("tentenct", "tentenct")

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	// Flush writes
	writeAPI.Flush()
	// get non-blocking write client
	//writeAPI := client.WriteAPI("tentenct", "tentenct")
	//需要增加时间来处理，点的信息需要时间和值
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
	// Get query client
	queryAPI := client.QueryAPI("tentenct")

	query := `from(bucket:"tentenct")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Access data
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}
}

//13598297289
