package roman

import "strings"

func romanAdd(left string, right string) string {
	ones_plus_one := []struct{ base, result string }{}
	if right == "C" {
		ones_plus_one = baseIncreaseTable("C", "D", "M")
	} else if right == "X" {
		ones_plus_one = baseIncreaseTable("X", "L", "C")
	} else {
		ones_plus_one = baseIncreaseTable("I", "V", "X")
	}
	for _, pair := range ones_plus_one {
		if strings.Contains(left, pair.base) {
			return romanAdd(strings.Replace(left, pair.base, "", 1), pair.result)
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
