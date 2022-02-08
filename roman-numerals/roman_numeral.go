package main

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	// [0,1,2,3,4,5], len = 6
	for i := 0; i < len(roman); i++ {
		if objMapper[roman[i]] != 0 {
			if i+1 < len(roman) && objMapper[roman[i]] < objMapper[roman[i+1]] {
				total -= objMapper[roman[i]]
			} else {
				total += objMapper[roman[i]]
			}
		}
	}
	return total
}

type romanNumeral struct {
	Value  int
	Symbol string
}

var objMapper = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}
var allRomanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
