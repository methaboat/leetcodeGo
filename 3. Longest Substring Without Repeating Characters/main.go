package main

import (
	"fmt"
)

// func lengthOfLongestSubstring(s string) int {
// 	var result int
// 	var count int
// 	var word string

// 	for _, j := range s {
// 		if word == "" {
// 			word = word + string(j)
// 			continue
// 		}
// 		isDup := false
// 		for _, k := range word {
// 			if j == k {
// 				isDup = true
// 				break
// 			}
// 		}
// 		if isDup == false {
// 			word = word + string(j)
// 		} else {
// 			count = len(word)
// 			if count >= result {
// 				result = count
// 				count = 0
// 				word = string(j)
// 				isDup = true
// 			}
// 		}
// 	}
// 	count = len(word)
// 	if count >= result {
// 		result = count
// 	}
// 	return result
// }
func lengthOfLongestSubstring(s string) int {
	var result int
	slice := []rune{}

	//loop based on inputs
	for _, input := range s {
		//check if the slice contains a value
		for i, a := range slice {
			if a == input {
				//extract the values after the corresponding value
				slice = slice[i+1:]
				break
			}
		}
		//add value
		slice = append(slice, input)
		// if the length of the slice is greater than the result, set the length
		if len(slice) > result {
			result = len(slice)
		}
	}

	return result
}
func main() {
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring("c"))
	// fmt.Println(lengthOfLongestSubstring(" "))
	// fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))

}
