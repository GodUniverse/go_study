package main

import (
	"fmt"
	"strings"
)

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

// 有效的括号
func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pair[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func main() {
	//题目一
	arr := []int{1, 1, 2, 2, 3, 4, 3}
	number := singleNumber(arr)
	fmt.Println(number)

	//题目二
	palindrome := isPalindrome(12321)
	fmt.Println(palindrome)

	//题目三
	valid := isValid("()")
	fmt.Println(valid)

	//题目四
	strs := []string{"hello", "hero", "head"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}
