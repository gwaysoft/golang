package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())

	t := now.Unix()

	fmt.Println(t, now.UnixNano())

	var n int8 = 1
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(time.Unix(t, 0))

	fmt.Println(now.Format("2006-01-02 15:04:05"), now.Add(time.Hour).Format("2006-01-02 15:04:05"))

	for tm := range time.Tick(time.Duration(30) * time.Second) {
		fmt.Println(tm)
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	tt, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-11-02 15:04:05", loc)
	fmt.Println(tt, err)
}
