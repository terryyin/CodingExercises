package roman

import "fmt"

func ExampleIsAddI() {
	fmt.Println(romanAdd("I", "I"))
	fmt.Println(romanAdd("II", "I"))
	fmt.Println(romanAdd("III", "I"))
	fmt.Println(romanAdd("IV", "I"))
	fmt.Println(romanAdd("V", "I"))
	fmt.Println(romanAdd("VI", "I"))
	fmt.Println(romanAdd("VII", "I"))
	fmt.Println(romanAdd("VIII", "I"))
	fmt.Println(romanAdd("IX", "I"))
	// Output:
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

func ExampleXsAddI() {
	fmt.Println(romanAdd("X", "I"))
	// Output:
	// XI
}
