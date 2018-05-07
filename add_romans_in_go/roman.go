package roman

import "strings"

func ToRoman(number int) string {
	romanDigit := map[int]string{
		4:  "IV",
		5:  "V",
		9:  "IX",
		10: "X",
	}
	if val, ok := romanDigit[number]; ok {
		return val
	}

	if number > 5 {
		return romanDigit[5] + ToRoman(number-5)
	}
	return strings.Repeat("I", number)
}
