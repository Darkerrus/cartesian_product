package main

import (
	"fmt"
	"strconv"

	"github.com/schwarmco/go-cartesian-product"
)

func main() {
	var res map[int]string = map[int]string{}
	var equation map[int][]string = map[int][]string{}
	str := "9876543210"
	op := []string{"-", "+", ""}
	var arr map[int][]interface{} = map[int][]interface{}{}

	for i := 0; i < len(str)-1; i++ {
		value := string(str[i])
		for j := 0; j < 3; j++ {
			v := op[j] + value
			arr[i] = append(arr[i], v)
		}
	}
	arr[9] = append(arr[9], "0")
	c := cartesian.Iter(arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9])
	var a int
	for product := range c {
		a++
		for i := 0; i < len(product); i++ {
			str := fmt.Sprintf("%v", product[i])
			res[a] += str
		}
	}
	for i := 0; i < len(res); i++ {
		arr := res[i]
		str := ""
		for j := 0; j < len(arr); j++ {
			if arr[0] == '+' {
				// delete(res, i)
			} else {
				if j == 0 {
					str += string(arr[j])
				} else {
					if arr[j] != '+' && arr[j] != '-' {
						str += string(arr[j])
					} else {
						equation[i] = append(equation[i], str)
						str = string(arr[j])
					}
				}
			}

		}
		equation[i] = append(equation[i], str)
	}
	for i := 0; i < len(equation); i++ {
		arr := equation[i]
		res := 0
		for j := 0; j < len(arr); j++ {
			value, _ := strconv.Atoi(arr[j])
			res += value
			if j == (len(arr) - 1) {
				if res == 200 {
					fmt.Println(equation[i])
				}
			}
		}
	}

}
