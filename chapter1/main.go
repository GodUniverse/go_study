package main

import "fmt"

// 只出现一次的数字
func singleNumber(nums []int) int {

	map1 := map[int]int{}

	for _, v := range nums {
		map1[v] = map1[v] + 1
	}
	for k, v := range map1 {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 回文数
func isPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	// 个位数是回文数
	if x < 10 {
		return true
	}

	//把整数转换为字符串
	str := fmt.Sprintf("%d", x)
	// 从字符串的两端开始比较字符
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//题目一
	arr := []int{1, 1, 2, 2, 3, 4, 3}
	number := singleNumber(arr)
	fmt.Println(number)

	//题目二
	palindrome := isPalindrome(12321)
	fmt.Println(palindrome)
}
