// Calculates my age in weeks
// Verified with https://www.timeanddate.com/date/timezoneduration.html
package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

const (
	// NOTE: PST doesn't parse (was still -0000 despite displaying PST)
	// GitHub Issue: https://github.com/golang/go/issues/24071
	// "It is not a goal that time.Time.Format and time.Parse be exact reverses of each other."
	// Time zone abbreviations such as "CST" are ambiguous: https://en.wikipedia.org/wiki/List_of_time_zone_abbreviations
	layout     = "2006-01-02 3:04 PM"
	dateFormat = "Monday, January 2, 2006 at 3:04 PM (MST)"
)

// Parse date/time with IANA Time Zone
func parseTime(date, zone string) (time.Time, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(layout, date, loc)
}

// Convert duration to an integer number of weeks, days, hours, and minutes
func convertDuration(duration time.Duration) (weeks, days, hours, minutes int) {
	weeksF := duration.Hours() / 24 / 7
	daysF := (weeksF - math.Trunc(weeksF)) * 7
	hoursF := (daysF - math.Trunc(daysF)) * 24
	minutesF := (hoursF - math.Trunc(hoursF)) * 60

	weeks = int(math.Trunc(weeksF))
	days = int(math.Trunc(daysF))
	hours = int(math.Trunc(hoursF))
	minutes = int(minutesF)

	return weeks, days, hours, minutes
}

func main() {
	// NOTE: PST because daylight saving time started in B.C. on Sunday, April 24, 1977
	birthdate, err := parseTime("1977-04-05 11:58 AM", "America/Vancouver")

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("The current time is %v\n\n", time.Now().Format(dateFormat))

	fmt.Printf("Nathan was born on %v\n", birthdate.Format(dateFormat))
	duration := time.Since(birthdate)

	weeks, days, hours, minutes := convertDuration(duration)
	fmt.Printf("He has been alive for %d weeks, %d days, %d hours and %d minutes.\n", weeks, days, hours, minutes)
}
