package roman

import "fmt"

func ExampleAddI() {
	fmt.Println(romanAdd("I", "I"))
	fmt.Println(romanAdd("II", "I"))
	fmt.Println(romanAdd("III", "I"))
	// Output:
	// II
	// III
	// IV
}
