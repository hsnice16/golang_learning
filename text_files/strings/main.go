package main

import (
	"fmt"
	"strconv"
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

	fmt.Println("\n----------------- strconv - Append Series -----------------")

	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println("string(str) -", string(str))

	fmt.Println("\n----------------- strconv - Format Series -----------------")

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println("a, b, c, d, e -", a, b, c, d, e)

	fmt.Println("\n----------------- strconv - Parse Series -----------------")

	boolA, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}

	floatB, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}

	intC, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	uIntD, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	intE := strconv.Itoa(1023)
	fmt.Println("boolA, floatB, intC, uIntD, intE -", boolA, floatB, intC, uIntD, intE)
}
