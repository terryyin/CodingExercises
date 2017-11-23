package roman

func romanAdd(left string, right string) string {
	ones := []string{"II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	ones_plus_one := map[string]string{}
	key := "I"
	for _, val := range ones {
		ones_plus_one[key] = val
		key = val
	}

	return ones_plus_one[left]
}
