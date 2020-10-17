package main

import (
	"flag"
	"fmt"
)

func main() {

	csvFilename := flag.String("filename", "algo por defecto", "enter a filename (must be .csv)")
	flag.Parse()
	fmt.Println(*csvFilename)
}
