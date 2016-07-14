package main

import (
    "fmt"
    w "github.com/zayanh/saigo/exercise-001-corpus/WordCount"
)


func main() {
    var testEntries w.Entries
    for i := 0; i < 10; i++ {
        testEntries = append(testEntries, &w.Entry{Word: string(byte(0x61+i)), Count: i})
    }
    b := testEntries.Less(1,0)
    fmt.Println(b)

    // Does it make sense to pass in a file handle instead of a string?
    _, entries := w.WordCount("7oldsamr.txt")

    // Can add some error checking here based on result

    // Print out the result of WordCount
    for _,val := range entries {
        fmt.Println(val.Count, val.Word)
    }
}