package main

import "fmt"

func titleToNumber(columnTitle string) int {
	var result int
	for _, j := range columnTitle {
		result = result * 26
		result += (int(j-'A') + 1)
		//fmt.Println(result)
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}
