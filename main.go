package main

import (
	"fmt"
	"strconv"
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
	c := cartesianProduct(arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9])
	var a int
	for _, v := range c {

		a++
		for i := 0; i < len(v); i++ {
			str := fmt.Sprintf("%v", v[i])
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

func cartesianProduct(arrs ...[]interface{}) [][]interface{} {
	if len(arrs) == 0 {
		return [][]interface{}{}
	}
	if len(arrs) == 1 {
		result := [][]interface{}{}
		for _, val := range arrs[0] {
			result = append(result, []interface{}{val})
		}
		return result
	}

	head, tail := arrs[0], arrs[1:]
	tailCartesian := cartesianProduct(tail...)

	result := [][]interface{}{}
	for _, val := range head {
		for _, subArr := range tailCartesian {
			result = append(result, append([]interface{}{val}, subArr...))
		}
	}

	return result
}
