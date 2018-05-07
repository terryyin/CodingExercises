package roman

import "fmt"

func ExampleOfRomanNumbers() {
	for number := 1; number <= 10; number++ {
		fmt.Println(ToRoman(number))
	}
	// Output:
	// I
	// II
	// III
	// IV
	// V
	// VI
	// VII
	// VIII
	// IX
	// X
}
