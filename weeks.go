// Calculates my age in weeks
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
	parseLayout = "2006-01-02 3:04 PM"
	dateFormat  = "Monday, January 2, 2006 at 3:04 PM (MST)"
)

func main() {
	birth, err := parseTime(birthTime, birthZone)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	weeks, days, hours, minutes := splitDuration(time.Since(birth))

	fmt.Printf("The current time is %v\n\n", time.Now().Format(dateFormat))
	fmt.Printf("%v was born on %v\n", name, birth.Format(dateFormat))
	fmt.Printf("%v has been alive for %.f weeks, %.f days, %.f hours and %.f minutes.\n", pronoun, weeks, days, hours, minutes)
}

// Parse date/time with IANA Time Zone
func parseTime(dateTime, zone string) (time.Time, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(parseLayout, dateTime, loc)
}

// Split duration into weeks, days, hours, and minutes
func splitDuration(duration time.Duration) (weeks, days, hours, minutes float64) {
	minutes = duration.Minutes()
	weeks, minutes = divMod(minutes, 7*24*60)
	days, minutes = divMod(minutes, 24*60)
	hours, minutes = divMod(minutes, 60)
	return weeks, days, hours, minutes
}

// divMod divides x/y and returns the quotient and remainder
func divMod(x, y float64) (float64, float64) {
	// NOTE: math.Mod to return the remainder, not math.Remainder
	return math.Trunc(x / y), math.Mod(x, y)
}
