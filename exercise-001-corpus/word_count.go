package main

import (
    "fmt"
    "os"

    "github.com/zayanh/saigo/exercise-001-corpus/corpus"
)


func main() {
    if len(os.Args) != 2 {
        fmt.Println("Wrong number of arguments, needs filename")
        os.Exit(1)
    }

    filename := os.Args[1]
    entries, err := corpus.WordCount(filename)
    if err != nil {
        fmt.Println("there was an error: " + err.Error())
        os.Exit(1)
    }

    // Print out the result of WordCount
    for _, val := range entries {
        fmt.Println(val.Count, val.Word)
    }
}