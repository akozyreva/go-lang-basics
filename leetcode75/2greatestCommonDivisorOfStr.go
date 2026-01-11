// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t
// (i.e., t is concatenated with itself one or more times).
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
package main

import (
	"fmt"
	"strings"
)

func getDivisorOfStr(word string) string {
	divisor := word[:2]
	isDivisorFound := false
	for !isDivisorFound {
		countOfOccur := strings.Count(word, divisor)
		if countOfOccur*len(divisor) == len(word) {
			isDivisorFound = true
		} else {
			index := strings.Index(word, divisor)
			divisor += string(word[index+len(divisor)])
		}
	}
	return divisor
}

func gcdOfStrings(str1 string, str2 string) string {
	divisor := str1[:2]
	if str2[:2] != divisor {
		return ""
	}
	divisor1 := getDivisorOfStr(str1)
	divisor2 := getDivisorOfStr(str2)
	if divisor1 == divisor2 {
		return divisor1
	} else {
		return ""
	}
}

// Claude decision:
func gcdOfStrings1(str1, str2 string) string {
	// Быстрая проверка: если конкатенация в разном порядке не равна - нет общего делителя
	if str1+str2 != str2+str1 {
		return ""
	}

	// Длина делителя = НОД длин строк
	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

// Наибольший общий делитель (алгоритм Евклида)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	divisor := gcdOfStrings("ABCABC", "ABC")
	fmt.Println(divisor)

	divisor = gcdOfStrings("LEET", "CODE")
	fmt.Println(divisor)

	divisor = gcdOfStrings("ABABAB", "ABAB")
	fmt.Println(divisor)

	divisor = gcdOfStrings("AAAAAB", "AAA")
	fmt.Println(divisor)

	divisor = gcdOfStrings("ABCABC", "AB")
	fmt.Println(divisor)
}
