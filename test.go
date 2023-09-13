package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type temp struct {
	Mess  string `json:"mess"`
	Oline string `json:"oline"`
}

type test1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Type string `json:"type"`
	Data temp   `json:"data"`
}

type test2 struct {
	userName  string
	userAge   int
	userType  string
	userMess  string
}

func test(start time.Time, function func(data string), data string) float64 {
	function(data)
	dtime := time.Since(start).Seconds()
	return dtime
}

func main() {

	testString1 := []string{}
	testString2 := []string{}
	test1Time := float64(0)
	test2Time := float64(0)
	num := 100000
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			testString1 = append(testString1, "{\"name\":\"test\",\"age\":150,\"type\":\"test\",\"data\":{\"mess\":\"test\",\"oline\":\"test\"}}")
		}
		if i%2 == 1 {
			testString1 = append(testString1, "{\"name\":\"test\",\"age\":8051,\"type\":\"test\",\"data\":{\"mess\":\"test\",\"oline\":\"test\"}}")
		}
	}
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			testString2 = append(testString2, "{\"name\":\"test\",\"age\":8694,\"type\":\"test\",\"data\":{\"mess\":\"test\",\"oline\":\"test\"}}")
		}
		if i%2 == 1 {
			testString2 = append(testString2, "{\"name\":\"test\",\"age\":452,\"type\":\"test\",\"data\":{\"mess\":\"test\",\"oline\":\"test\"}}")
		}
	}
	for i := 0; i < num; i++ {
		test1Time += test(time.Now(), func(data string) {
			var test test1
			var test2 test2
			json.Unmarshal([]byte(data), &test)
			test2.userName = test.Name
			test2.userAge = test.Age
			test2.userType = test.Type
			test2.userMess = test.Data.Mess
		}, testString1[i])
	}
	reg1 := regexp.MustCompile(`"name":"(.+?)"`)
	reg2 := regexp.MustCompile(`"age":([0-9]+)`)
	reg3 := regexp.MustCompile(`"type":"(.+?)"`)
	reg4 := regexp.MustCompile(`"mess":"(.+?)"`)
	for i := 0; i < num; i++ {
		test2Time += test(time.Now(), func(data string) {
			var test2 test2
			test2.userName = reg1.FindStringSubmatch(data)[1]
			test2.userAge, _ = strconv.Atoi(reg2.FindStringSubmatch(data)[1])
			test2.userType = reg3.FindStringSubmatch(data)[1]
			test2.userMess = reg4.FindStringSubmatch(data)[1]
		}, testString2[i])
	}
	fmt.Println("json.Unmarshal:  ", test1Time, "s")
	fmt.Println("regexp:  ", test2Time, "s")

}
