package main

import (
	"fmt"
	"os"
	"regexp"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}

	return true
}

func main() {
	isIp := IsIP("127.0.0.1")
	fmt.Println("Is IP:", isIp)

	isIp = IsIP("127.0.0")
	fmt.Println("Is IP:", isIp)

	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("Number")
	} else {
		fmt.Println("Not number")
	}
}
