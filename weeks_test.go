package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSplitDuration(t *testing.T) {
	var tests = []struct {
		duration                    string
		weeks, days, hours, minutes float64
	}{
		{duration: "0m", weeks: 0, days: 0, hours: 0, minutes: 0},
		{duration: "1h10m", weeks: 0, days: 0, hours: 1, minutes: 10},
		{duration: "746h", weeks: 4, days: 3, hours: 2, minutes: 0},
	}
	for _, test := range tests {
		duration, _ := time.ParseDuration(test.duration)
		weeks, days, hours, minutes := splitDuration(duration)
		assertInDelta(t, fmt.Sprintf("%v to have weeks", test.duration), test.weeks, weeks, epsilon)
		assertInDelta(t, fmt.Sprintf("%v to have days", test.duration), test.days, days, epsilon)
		assertInDelta(t, fmt.Sprintf("%v to have hours", test.duration), test.hours, hours, epsilon)
		assertInDelta(t, fmt.Sprintf("%v to have minutes", test.duration), test.minutes, minutes, epsilon)
	}
}

const epsilon = 1e-9

func assertInDelta(t *testing.T, msg string, expected, actual, delta float64) {
	if math.Abs(expected-actual) > delta {
		t.Errorf("Expected %v %.f, but got %.f", msg, expected, actual)
	}
}
