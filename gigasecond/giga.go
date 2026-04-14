package main

import (
	"fmt"
	"time"
)

func gigasecond(t time.Time) time.Time {
	GIGA := 1_000_000_000

	gt := t.Add(time.Second * time.Duration(GIGA))
	return gt
}

func main() {
	var day, month, year int
	fmt.Print("What is the day(as an int): ")
	fmt.Scan(&day)
	fmt.Print("What is the month(as in int): ")
	fmt.Scan(&month)
	fmt.Print("What is the year(as an int): ")
	fmt.Scan(&year)
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	nt := gigasecond(t)

	fmt.Printf("The gigasecond from the date you entered is %v", nt)
}
