package main

import (
	"fmt"
	"time" // jangan lupa import package time
)

func main() {
	var timeDateLocal time.Time = time.Date(2022, 1, 16, 10, 15, 0, 0, time.Local)
	var timeDateUTC time.Time = time.Date(2022, 1, 16, 10, 15, 0, 0, time.UTC)
	var timeNow time.Time = time.Now()

	fmt.Printf("Time Date Local: %v\n", timeDateLocal)
	fmt.Printf("Time Date UTC: %v\n", timeDateUTC)
	fmt.Printf("Time Now: %v\n", timeNow)

	// Fungsi parse
	layout := "2006-01-02 15:04:05"
	myTimeStr := "2021-01-16 10:00:00"

	myTime, err := time.Parse(layout, myTimeStr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(myTime.String())
	}

	parsedTime, err := time.Parse(time.RFC3339, "2022-01-16T15:04:05Z")
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("Oke", parsedTime)
	}

	// Method time.Time
	now := time.Now()
	fmt.Println("Waktu saat ini:", now.String())

	year := now.Year()
	month := now.Month()
	day := now.Day()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println("Tanggal:", day, month, year)
	fmt.Printf("Jam: %d:%d:%d\n", hour, minute, second)
}
