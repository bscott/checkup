package main

import (
	"fmt"
	"log"

	"bitbucket.org/mattholt/checkup"
)

func main() {
	c := checkup.Checkup{
		Checkers: []checkup.Checker{
			checkup.HTTPChecker{Name: "IP Chicken", URL: "http://ipchicken.com", Attempts: 5},
			checkup.HTTPChecker{Name: "Example", URL: "https://example.com", Attempts: 5},
		},
		Storage: checkup.S3{
			AccessKeyID:     "...",
			SecretAccessKey: "...",
			Region:          "us-east-1",
			Bucket:          "srcgraph-monitor-test",
			CheckExpiry:     24 * time.Hour * 7,
		},
	}

	results, err := c.Check()
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		stats := result.ComputeStats()
		fmt.Printf("== %s - %s\n", result.Title, result.Endpoint)
		fmt.Printf("        Max: %s\n", stats.Max)
		fmt.Printf("        Min: %s\n", stats.Min)
		fmt.Printf("     Median: %s\n", stats.Median)
		fmt.Printf("       Mean: %s\n", stats.Mean)
		fmt.Printf("        All: %v\n", result.Times)
		fmt.Printf(" Assessment: %v\n\n", result.Status())
	}
}