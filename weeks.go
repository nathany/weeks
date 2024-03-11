// Calculates my age in weeks
// Verified with https://www.timeanddate.com/date/timezoneduration.html
package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

// Person struct
type Person struct {
	name      string
	pronoun   string
	birthTime string
	birthZone string
}

const (
	// NOTE: PST doesn't parse (was still -0000 despite displaying PST)
	// GitHub Issue: https://github.com/golang/go/issues/24071
	// "It is not a goal that time.Time.Format and time.Parse be exact reverses of each other."
	// Time zone abbreviations such as "CST" are ambiguous: https://en.wikipedia.org/wiki/List_of_time_zone_abbreviations
	parseLayout = "2006-01-02 3:04 PM"
	dateFormat  = "Monday, January 2, 2006 at 3:04 PM (MST)"
)

func main() {
	var person = Person{name: "Nathan", pronoun: "He", birthTime: "1977-04-05 11:58 AM", birthZone: "America/Vancouver"}
	birthdate, err := parseTime(person.birthTime, person.birthZone)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("The current time is %v\n\n", time.Now().Format(dateFormat))

	// NOTE: PST because daylight saving time started in B.C. on Sunday, April 24, 1977
	fmt.Printf("%v was born on %v\n", person.name, birthdate.Format(dateFormat))

	duration := time.Since(birthdate)
	weeks, days, hours, minutes := convertDuration(duration)
	fmt.Printf("%v has been alive for %.f weeks, %.f days, %.f hours and %.f minutes.\n", person.pronoun, weeks, days, hours, minutes)
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
