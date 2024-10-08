package main

import (
	"fmt"
	"time"
)

func main () {

	timeNow :=  time.Now()
	fmt.Println("Time Now: ", timeNow)
	formattedTimeNowByDate := timeNow.Format("2006/01/02")
	
	fmt.Println("Formatted Time Now By Date: ", formattedTimeNowByDate)
	fmt.Println("Formatted Time by date: ", formattedTimeNowByDate)

	formattedTimeByDateAndTime := timeNow.Format("2006/01/02 Monday")
	fmt.Println("Formatted Time Now By Date And Time: ", formattedTimeByDateAndTime)

	formattedTime := timeNow.Format("15:04 PM")
	fmt.Println("Formatted Time: ", formattedTime)

}