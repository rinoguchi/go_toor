package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	printOs()
	printHour()
}

func printOs() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func printHour() {
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println(hour, "AM")
	default:
		fmt.Println(hour-12, "PM")
	}

}
