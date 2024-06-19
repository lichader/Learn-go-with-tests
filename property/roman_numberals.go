package property

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

// order matters?
var allRomanNumerals = []RomanNumeral{
	{100, "C"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(input int) string {
	var result strings.Builder

	for _, numberal := range allRomanNumerals {
		for input >= numberal.Value {
			result.WriteString(numberal.Symbol)
			input -= numberal.Value
		}
	}

	return result.String()
}
