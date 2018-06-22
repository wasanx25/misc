package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	duration := "PT1H1S"
	base := regexp.MustCompile(`[0-9]{1,}`)
	list := base.FindAllStringSubmatch(duration, -1)

	var timeList []int
	for _, time := range list {
		num, _ := strconv.Atoi(time[0])
		timeList = append(timeList, num)
	}

	var result []int
	if strings.Index(duration, "H") == -1 && strings.Index(duration, "M") >= 0 && strings.Index(duration, "S") == -1 { // 1M
		result = append(result, 0)
		result = append(result, timeList[0])
		result = append(result, 0)
	} else if strings.Index(duration, "H") >= 0 && strings.Index(duration, "M") == -1 && strings.Index(duration, "S") >= 0 { // 1H1S
		result = append(result, timeList[0])
		result = append(result, 0)
		result = append(result, timeList[1])
	} else if strings.Index(duration, "H") >= 0 && strings.Index(duration, "M") == -1 && strings.Index(duration, "S") == -1 { // 1H
		result = append(result, timeList[0])
		result = append(result, 0)
		result = append(result, 0)
	} else if strings.Index(duration, "H") >= 0 && strings.Index(duration, "M") >= 0 && strings.Index(duration, "S") == -1 { // 1H1M
		result = append(result, timeList[0])
		result = append(result, timeList[1])
		result = append(result, 0)
	}

	if len(result) == 0 {
		result = timeList
	}

	var seconds int
	switch len(result) {
	case 1:
		seconds = result[0]
	case 2:
		seconds = result[0]*60 + result[1]
	case 3:
		seconds = result[0]*60*60 + result[1]*60 + result[2]
	}

	fmt.Println(result)
	fmt.Println(seconds)

	// fmt.Println(duration)
	// duration = strings.TrimPrefix(duration, "PT")
	// fmt.Println(duration)
	// hoursIndex := strings.Index(duration, "H")
	// if hoursIndex != -1 {
	// 	hour := duration[0:hoursIndex]
	// 	duration = strings.TrimPrefix(duration, string(hour)+"H")
	// 	fmt.Println(hour)
	// }
	// var a int
	// fmt.Scan(&a)
	// fmt.Println(a)

	// fmt.Println(hour)
	// fmt.Println(duration[0:hour])
}
