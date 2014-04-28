package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: aren source_file target_file")
		os.Exit(1)
	}

	err := atomic_rename(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
