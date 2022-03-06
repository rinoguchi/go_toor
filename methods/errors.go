package main

import (
	"fmt"
	"time"
)

type MyError struct {
	Time   time.Time
	Reason string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s: %v", e.Reason, e.Time)
}

func calsSum(a int, b int) (int, error) {
	sum := a + b
	if sum > 100 {
		return 0, &MyError{time.Now(), "limit over"}
	}
	return sum, nil
}

func main() {
	var err error
	var sum int

	fmt.Println("========== normal case ==========")
	sum, err = calsSum(10, 80)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)

	fmt.Println("========== error case ==========")
	sum, err = calsSum(50, 80)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)
}
