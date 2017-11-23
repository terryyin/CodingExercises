package roman

import "strings"

func romanAdd(left string, right string) string {
	if right == "" {
		return left
	}
	if right == "IV" {
		left = romanAdd(left, right[0:len(right)-2]+"III")
		right = "I"
	} else if right == "V" {
		left = romanAdd(left, right[0:len(right)-1]+"IV")
		right = "I"
	} else {
		left = romanAdd(left, right[0:len(right)-1])
		right = "I"
	}

	return romanIncrease(left, right)
}

func romanIncrease(left string, right string) string {
	ones_plus_one := []struct{ base, result string }{}
	if right == "C" {
		ones_plus_one = baseIncreaseTable("C", "D", "M")
	} else if right == "X" {
		ones_plus_one = baseIncreaseTable("X", "L", "C")
	} else if right == "I" {
		ones_plus_one = baseIncreaseTable("I", "V", "X")
	}
	for _, pair := range ones_plus_one {
		if strings.Contains(left, pair.base) {
			return romanIncrease(strings.Replace(left, pair.base, "", 1), pair.result)
		}
	}
	return left + right
}

func baseIncreaseTable(single string, five string, ten string) []struct{ base, result string } {
	return []struct{ base, result string }{
		{single + ten, ten},
		{five + single + single + single, single + ten},
		{single + five, five},
		{single + single + single, single + five}}
}
