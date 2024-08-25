package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("hsnice16", 0777)
	os.MkdirAll("hsnice16/test1/test2", 0777)
	err := os.Remove("hsnice16")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("hsnice16")
}
