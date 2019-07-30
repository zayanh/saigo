package main

import (
	"fmt"
	w "github.com/zayanh/saigo/src/words"
	"os"
)

func main() {
	// define command line parameter flags
	params := os.Args
	if len(params) < 2 {
		fmt.Println("no input file provided. Exiting")
		return
	}
	filename := params[1]
	//TODO: Make sure path is valid

	fmt.Println(w.Count(filename))
}
