package main

import (
	"fmt"
	"leetcode-go.git/code1"
)

func main() {
	var nums = []int{3,2,4}
	var target int = 6
	a := code1.TwoSum(nums,target)
	fmt.Println(a)
}
