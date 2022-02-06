package main

import (
	"fmt"
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
}
