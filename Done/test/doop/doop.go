package main

import (
	"os"

	"github.com/01-edu/z01"
)

func checknumber(str string) bool {
	check := []rune(str)
	for _, x := range check {
		if check[0] == '-' {
			continue
		}
		if x < '0' || x > '9' {
			return false
		}
	}
	return true
}
func checkoperator(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "%" || str == "/" {
		return true
	} else {
		return false
	}

}
func stringtonumber(str string) int64 {
	check := []rune(str)
	ans, ok := 0
	start := 0
	if check[0] == '-' {
		start = 1
	}
	ln := 0
	for range check {
		ln++
	}
	for i := start; i < ln; i++ {
		ans *= 10
		ans = ans + int64(check[i]) - 48

	}
	if check[0] == '-' {
		return -ans
	}
	return ans
}
func printnumb(numb int64) {
	str := ""
	if numb == 0 {
		str = "0"
	}
	if numb < 0 {
		z01.PrintRune('-')
		numb *= -1
	}
	for numb != 0 {
		str = str + (string(numb%10 + 48))
		numb /= 10
	}
	ln := 0
	ptr := []rune(str)
	for range ptr {
		ln++
	}
	for j := ln - 1; j >= 0; j-- {
		z01.PrintRune(ptr[j])
	}
	z01.PrintRune('\n')
}
func printmessage(str string) {
	ptr := []rune(str)
	for _, p := range ptr {
		z01.PrintRune(p)
	}
	z01.PrintRune('\n')
}
func main() {
	ln := 0
	for range os.Args {
		ln++
	}

	// ln = 4 , value , operator, value
	if ln == 4 {
		value := os.Args[1]
		operator := os.Args[2]
		vl2 := os.Args[3]
		if checknumber(value) == true && checknumber(vl2) == true && checkoperator(operator) == true {
			v1, ok := stringtonumber(value)
			v2, ok := stringtonumber(vl2)
			if v1 < -9223372036854775807 || v1 > 9223372036854775807 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			if v2 < -9223372036854775807 || v2 > 9223372036854775807 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			var ans, ok = 0
			if operator == "+" {
				ans = v1 + v2
			} else if operator == "*" {
				ans = v1 * v2
			} else if operator == "-" {
				ans = v1 - v2
			} else if operator == "/" {
				if v2 == 0 {
					printmessage("No division by 0")
					return
				} else {
					ans = v1 / v2
				}
			} else if operator == "%" {
				if v2 == 0 {
					printmessage("No Modulo by 0")
					return
				} else {
					ans = v1 % v2
				}
			} else {

			}
			if ans < -9223372036854775808 || ans > 9223372036854775807 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
			} else {
				printnumb(ans)
			}
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		}
	}
}
