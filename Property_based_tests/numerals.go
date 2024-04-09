package numerals

import "strings"

// ConvertToRoman converts an Arabic numeral to a Roman numeral.
func ConvertToRoman(arabic int) string {
	var result strings.Builder
	var numeralMap = []struct {
		value  int
		symbol string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, numeral := range numeralMap {
		for arabic >= numeral.value {
			result.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}
	return result.String()
}

// ConvertToArabic converts a Roman numeral to an Arabic numeral.
func ConvertToArabic(roman string) int {
	var numeralMap = map[string]int{
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	var total int
	for i := 0; i < len(roman); {
		if i+1 < len(roman) && numeralMap[roman[i:i+2]] > 0 {
			total += numeralMap[roman[i:i+2]]
			i += 2
		} else {
			total += numeralMap[string(roman[i])]
			i++
		}
	}
	return total
}
