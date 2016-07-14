package main

import (
    "fmt"
    "os"
    w "github.com/zayanh/saigo/exercise-001-corpus/WordCount"
)


func main() {
    _, entries := w.WordCount(string(os.Args[1]))

    // Can add some error checking here based on result

    // Print out the result of WordCount
    for _,val := range entries {
        fmt.Println(val.Count, val.Word)
    }
}