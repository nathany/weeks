// Calculates my age in weeks
// Verified with https://www.timeanddate.com/date/timezoneduration.html
package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

// Me
const (
	name      = "Nathan"
	pronoun   = "He"
	birthTime = "1977-04-05 11:58 AM"
	birthZone = "America/Vancouver"
)

// Formats to parse and display times
const (
	// Time zone abbreviations such as "CST" are ambiguous: https://en.wikipedia.org/wiki/List_of_time_zone_abbreviations
	// So we can't use MST in the parseLayout even though it is available for displaying the timezone.
	// GitHub Issue: https://github.com/golang/go/issues/24071
	// "It is not a goal that time.Time.Format and time.Parse be exact reverses of each other."
	parseLayout = "2006-01-02 3:04 PM"
	dateFormat  = "Monday, January 2, 2006 at 3:04 PM (MST)"
)

func main() {
	// NOTE: Returns time in PST because daylight saving time started in B.C. on Sunday, April 24, 1977
	when, err := parseTime(birthTime, birthZone)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	weeks, days, hours, minutes := convertDuration(time.Since(when))

	fmt.Printf("The current time is %v\n\n", time.Now().Format(dateFormat))
	fmt.Printf("%v was born on %v\n", name, when.Format(dateFormat))
	fmt.Printf("%v has been alive for %.f weeks, %.f days, %.f hours and %.f minutes.\n", pronoun, weeks, days, hours, minutes)
}

// Parse date/time with IANA Time Zone
func parseTime(date, zone string) (time.Time, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(parseLayout, date, loc)
}

// Convert duration to an integer number of weeks, days, hours, and minutes
func convertDuration(duration time.Duration) (weeks, days, hours, minutes float64) {
	minutes = duration.Minutes()
	weeks, minutes = divMod(minutes, 24*7*60)
	days, minutes = divMod(minutes, 24*60)
	hours, minutes = divMod(minutes, 60)

	return weeks, days, hours, minutes
}

// divMod divides x/y and returns the quotient and remainder
func divMod(x, y float64) (float64, float64) {
	return math.Trunc(x / y), math.Mod(x, y)
}
