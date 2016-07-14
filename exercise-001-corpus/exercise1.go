package main

import (
    "fmt"
    "io/ioutil"
    "sort"
    "strings"
)

type Entry struct {
    word  string
    count int
}

// Had to be pointers because the slice won't update
// properly otherwise
type Entries []*Entry

// The number of elements in the collection
func (e Entries) Len() int {
    return len(e)
}

// Reports if the element with index i should sort before
// the element with index j
func (e Entries) Less(i, j int) bool {
    if e[i].count > e[j].count {
        return true
    } else {
        return false
    }
}

// Swaps the elements with indexes i and j
func (e Entries) Swap(i, j int) {
    tmp := Entry{word: e[i].word, count: e[i].count}
    e[i] = e[j]
    e[j] = &tmp
}

func main() {
    // Open, read and put contents of the file in a string variable
    bytes, err := ioutil.ReadFile("7oldsamr.txt")
    // bytes, err := ioutil.ReadFile("7oldsamr.txt")
    if err != nil {
        fmt.Println("Error reading file")
        return
    }
    str := string(bytes)

    // Get rid of punctuation
    str = strings.Replace(str, "\"", "", -1)
    str = strings.Replace(str, ",", "", -1)
    str = strings.Replace(str, "!", "", -1)
    str = strings.Replace(str, ":", "", -1)
    str = strings.Replace(str, "-", "", -1)
    str = strings.Replace(str, ".", "", -1)
    str = strings.Replace(str, "    ", "", -1)

    // Replace newlines with single space
    str = strings.Replace(str, "\n", " ", -1)

    // Make it all lower case to compare
    str = strings.ToLower(str)

    // Split str into a slice of type string at every instance of " "
    words := strings.Fields(str)

    // A slice of type Entry
    var entries Entries

    // Iterate over each word
    for _, word := range words {

        // Iterate over all the words we have seen already
        exist := 0
        for _,e := range entries {
            if e.word == word {
                e.count++
                exist = 1
                break
            }
        }

        // Add a new entry in our list if this is a new word
        if exist == 0 {
            entries = append(entries, &Entry{word: word, count: 1})
        }
    }

    // Let's check our handywork
    sort.Sort(entries)
    for _,val := range entries {
        fmt.Println(val.count, val.word)
    }
}
