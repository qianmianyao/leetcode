package LeetCode

import (
	"strings"
)

// ContainsDuplicate 数组中有没有重复的元素
func ContainsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

// TwoSum 两数之和
// 假如 nums = [2,7,11,15], target = 9
// 那么先建立一个hash表，然后循环这个数组得到索引和值，在第一次循环的时候，9-2=7，查询hash表中没有这个键，那么久把2加入到hash表中
// 在第二次循环的时候，9-7=2，这次在hash表中查询到了2，就直接返回当前7的索引，和hash表中2的索引
func TwoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, ok := set[complement]; ok {
			return []int{index, i}
		}
		set[num] = i
	}
	return nil
}

// IsPalindrome 判断是否是回文数
// 假如 x=1221
func IsPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	if x == 0 {
		return true
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	if x == revertedNumber || x == revertedNumber/10 {
		return true
	}
	return false
}

// RomanToInt 罗马数字转整数
func RomanToInt(s string) int {
	s = strings.Replace(s, "IV", "a", -1)
	s = strings.Replace(s, "IX", "b", -1)
	s = strings.Replace(s, "XL", "c", -1)
	s = strings.Replace(s, "XC", "d", -1)
	s = strings.Replace(s, "CD", "e", -1)
	s = strings.Replace(s, "CM", "f", -1)

	result := 0
	for _, v := range strings.Split(s, "") {
		result += getValue(v)
	}
	return result
}

func getValue(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	case "a":
		return 4
	case "b":
		return 9
	case "c":
		return 40
	case "d":
		return 90
	case "e":
		return 400
	case "f":
		return 900
	}
	return 0
}

// LongestCommonPrefix 最长公共前缀
func LongestCommonPrefix(strList []string) string {
	if len(strList) == 0 {
		return ""
	}
	return ""
}
