package main

import "fmt"

func main() {
	stringFormat := "%10s\n"
	fmt.Printf(stringFormat, "*")

	var charAmount int
	fmt.Print("Insert char amount: ")
	fmt.Scan(&charAmount)

	// %[number]s -> %%[number]s (prints '% d (MISSING) s' because go understands double percent as a scape to use percent)
	// so in order to "scape the scape", use three percent signs: "%%%ds\n"
	dynamicStringFormat := fmt.Sprintf("%%%ds\n", charAmount)
	fmt.Printf(dynamicStringFormat, "*")

	// To align in the left use the pattern "%-[numberOfChars]"
}
