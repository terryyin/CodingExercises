package roman

import "strings"

func romanAdd(left string, right string) string {
	ones_plus_one := []struct{ base, result string }{{"IX", "X"}, {"VIII", "IX"}, {"IV", "V"}, {"III", "IV"}}

	for _, pair := range ones_plus_one {
		if strings.Contains(left, pair.base) {
			return strings.Replace(left, pair.base, "", 1) + pair.result
		}
	}

	return left + right
}
