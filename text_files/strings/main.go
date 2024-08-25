package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("----------------- Contains -----------------")

	fmt.Println(`strings.Contains("seafood", "foo") -`, strings.Contains("seafood", "foo"))
	fmt.Println(`strings.Contains("seafood", "bar") -`, strings.Contains("seafood", "bar"))
	fmt.Println(`strings.Contains("seafood", "") -`, strings.Contains("seafood", ""))
	fmt.Println(`strings.Contains("", "") -`, strings.Contains("", ""))

	fmt.Println("\n----------------- Join -----------------")

	s := []string{"foo", "bar", "baz"}
	fmt.Println(`strings.Join(s, ", ") -`, strings.Join(s, ", "))

	fmt.Println("\n----------------- Repeat -----------------")

	fmt.Println(`"ba" + strings.Repeat("na",2) -`, "ba"+strings.Repeat("na", 2))

	fmt.Println("\n----------------- Replace -----------------")

	fmt.Println(`strings.Replace("oink oink oink", "k", "ky", 2) -`, strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(`strings.Replace("oink oink oink", "oink", "moo", -1) -`, strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Println("\n----------------- Split -----------------")

	fmt.Printf(`strings.Split("a,b,c", ",") - %q`+"\n", strings.Split("a,b,c", ","))
	fmt.Printf(`strings.Split("a man a plan a canal panama", "a ") - %q`+"\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf(`strings.Split(" xyz ", "") - %q`+"\n", strings.Split(" xyz ", ""))
	fmt.Printf(`strings.Split("", "Bernardo O'Higgings") - %q`+"\n", strings.Split("", "Bernardo O'Higgings"))

	fmt.Println("\n----------------- Trim -----------------")

	fmt.Printf(`strings.Trim(" !!! Achtung !!! ", "! ") - [%q]`+"\n", strings.Trim(" !!! Achtung !!! ", "! "))

	fmt.Println("\n----------------- Fields -----------------")

	fmt.Printf(`strings.Fields("  foo bar baz  ") - %q`+"\n", strings.Fields("  foo bar baz  "))
}
