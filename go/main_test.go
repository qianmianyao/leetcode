//go:build testing

package LeetCode

import (
	"fmt"
	"testing"
)

func TestBasicGrammar(t *testing.T) {
	fmt.Println("Hello, world!")
}

func TestContainsDuplicate(t *testing.T) {
	num := []int{1, 2, 3, 1}
	results := ContainsDuplicate(num)
	fmt.Println(results)
}

func TestTwoSum(t *testing.T) {
	num := []int{3, 3}
	target := 6
	results := TwoSum(num, target)
	fmt.Println(results)
}

func TestIsPalindrome(t *testing.T) {
	x := 1122112211
	result := IsPalindrome(x)
	fmt.Println(result)
}

func TestRomanToInt(t *testing.T) {
	result := RomanToInt("MCMXCIV")
	fmt.Println(result)
}

func TestLongestCommonPrefix(t *testing.T) {
	strList := []string{"add", "acc", "abb", "asscd"}
	result := LongestCommonPrefix(strList)
	fmt.Println(result)
}
