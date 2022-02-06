package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ip validation
func Is_valid_ip(ip string) bool {
	arr := strings.Split(ip, ".")
	if len(arr) != 4 {
		return false
	}
	for _, v := range arr {
		if v[0] == '0' && len(v) > 1 {
			return false
		}
		i, err := strconv.Atoi(v)

		if err != nil {
			return false
		}

		if i < 0 || i > 255 {
			return false
		}
	}

	return true
}

// deadfish command
func Parse(data string) []int {
	result := []int{}
	temp := 0
	m1 := regexp.MustCompile(`[^idso]`)

	tempData := m1.ReplaceAll([]byte(data), []byte(""))
	data = string(tempData)

	// 	i increments the value (initially 0)
	// d decrements the value
	// s squares the value
	// o outputs the value into the return array

	l := len(data)
	for i := 0; i < l; i++ {
		switch data[i] {
		case 'i':
			temp += 1
		case 'd':
			temp -= 1
		case 's':
			temp = temp * temp
		case 'o':
			result = append(result, temp)
		default:
			continue
		}
	}

	return result
}

func WordsToMarks(s string) int {
	var sum int
	for _, v := range s {
		sum += int(v - 96)
	}
	return sum
}

func solve(str string) string {
	m1 := regexp.MustCompile(`[a-z]`)

	tempData := m1.ReplaceAll([]byte(str), []byte(""))
	data := string(tempData)

	if len(data) > (len(str) - len(data)) {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}

func main() {

	//ip validation
	fmt.Println(Is_valid_ip("1.2.3.4"))
	fmt.Println(Is_valid_ip("123.45.67.89"))

	fmt.Println(Is_valid_ip("1.2.3"))
	fmt.Println(Is_valid_ip("1.2.3.4.5"))
	fmt.Println(Is_valid_ip("123.456.78.90"))
	fmt.Println(Is_valid_ip("123.045.067.089"))

	fmt.Println(Is_valid_ip("12.34.56 .1"))
	fmt.Println(Is_valid_ip("123.045.067.089"))

	// deadhfish command
	fmt.Println(Parse("ssdizbsnsdkkincuybidds"))

	// word to sum
	fmt.Println(WordsToMarks("az"))

	fmt.Println(200 - 200%37)

	fmt.Println(
		solve("coDe"),
		solve("cODE"),
		solve("coDE"))

}
