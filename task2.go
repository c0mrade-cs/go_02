package main

import (
	"fmt"
	"time"
)

func main() {
	dobStr := "03.03.2000" // Replace this date with your birthday
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %s\n", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday string) {
	weekday = date.Weekday().String()
	return
}
