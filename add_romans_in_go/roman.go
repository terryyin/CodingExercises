package roman

func romanAdd(left string, right string) string {
	ones := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	ones_plus_one := map[string]string{}
	key := ""
	for _, val := range ones {
		ones_plus_one[key] = val
		key = val
	}

	if val, ok := ones_plus_one[left]; ok {
		return val
	}
	return left[0:1] + romanAdd(left[1:], right)
}
