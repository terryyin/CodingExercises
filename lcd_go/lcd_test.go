package lcd

import "fmt"

func lcd(time string) string {
	return "\n\n   |\n   |\n\n"
}

func Example1() {
	fmt.Print("---")
	fmt.Print(lcd("1"))
	fmt.Print("---")
	// Output:
	// ---
	//
	//    |
	//    |
	//
	// ---

}
