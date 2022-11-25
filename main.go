package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"os"
)

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	arg, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, arg)

	fmt.Println(string(contents))

}
