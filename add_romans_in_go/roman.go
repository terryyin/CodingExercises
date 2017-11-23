package roman

import "strings"

func romanAdd(left string, right string) string {
	ones_plus_one := []struct{ base, result string }{{"IX", "X"}, {"VIII", "IX"}, {"IV", "V"}, {"III", "IV"}}
	for _, pair := range ones_plus_one {
		if strings.Contains(left, pair.base) {
			return romanTensAdd(strings.Replace(left, pair.base, "", 1), pair.result)
		}
	}
	return left + right
}

func romanTensAdd(left string, right string) string {
	if right == "X" {
		ones_plus_one := []struct{ base, result string }{{"XC", "C"}, {"LXXX", "XC"}, {"XL", "L"}, {"XXX", "XL"}}
		for _, pair := range ones_plus_one {
			if strings.Contains(left, pair.base) {
				return strings.Replace(left, pair.base, "", 1) + pair.result
			}
		}
	}
	return left + right
}
