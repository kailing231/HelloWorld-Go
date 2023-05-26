package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// // Prints each index and arg
	// for i, arg := range os.Args {
	// 	fmt.Println(i, arg)
	// }
	// // the filename input is os.Args[1]

	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// os.Stdout is *File
	io.Copy(os.Stdout, fp)

	// build main.go, execute main.exe, doesnt work on Code's Terminal
}
