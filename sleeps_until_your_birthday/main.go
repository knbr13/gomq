package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02" // equivalent to 'time.DateOnly' constant

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf(
		"You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)),
	)
}

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	parsedTime, err := time.Parse(expectedFormat, target)
	if err != nil {
		log.Fatal("invalid target date: ", target)
	}
	return parsedTime
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	now := time.Now()
	target = target.AddDate(now.Year()-target.Year(), 0, 0)
	hours := time.Until(target).Hours()
	return float64(hours/24) + 1.0
}
