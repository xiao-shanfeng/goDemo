package main

import "fmt"

const (
	SecondsPerMinute = 60	//每分钟的秒数
	SecondsPerHour = SecondsPerMinute * 60	//每小时的秒数
	SecondsPerDay = SecondsPerHour * 24	//每天的秒数
)

func parseTime() {
	fmt.Println(resolveTime(1000))
	// 至获取小时和分钟
	_,hour,minute := resolveTime(18000)
	fmt.Println(hour,minute)
	day,_,_ := resolveTime(90000)
	fmt.Println(day)
}

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}