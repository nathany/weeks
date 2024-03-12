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

func TestDivMod(t *testing.T) {
	var tests = []struct {
		dividend, divisor, quotient, remainder float64
	}{
		{dividend: 0, divisor: 1, quotient: 0, remainder: 0},
		{dividend: 400, divisor: 10, quotient: 40, remainder: 0},
		{dividend: 300, divisor: 44, quotient: 6, remainder: 36},
	}
	for _, test := range tests {
		q, r := divMod(test.dividend, test.divisor)
		assertInDelta(t, fmt.Sprintf("%v/%v to have quotient", test.dividend, test.divisor), test.quotient, q, epsilon)
		assertInDelta(t, fmt.Sprintf("%v/%v to have remainder", test.dividend, test.divisor), test.remainder, r, epsilon)
	}
}

func TestDivModDivideByZero(t *testing.T) {
	q, r := divMod(14, 0)
	if !math.IsInf(q, 1) {
		t.Errorf("Expected %v/%v quotient to be +Inf, but got %.f", 14, 0, q)
	}
	if !math.IsNaN(r) {
		t.Errorf("Expected %v/%v remainder to be NaN, but got %.f", 14, 0, r)
	}
}

const epsilon = 1e-9

func assertInDelta(t *testing.T, msg string, expected, actual, delta float64) {
	if math.Abs(expected-actual) > delta {
		t.Errorf("Expected %v %.f, but got %.f", msg, expected, actual)
	}
}
