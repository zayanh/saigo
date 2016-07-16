package corpus

import (
    // "fmt"
    "io/ioutil"
    "sort"
    "strings"
)

type Entry struct {
    Word  string
    Count int
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
    return e[i].Count > e[j].Count
}

// Swaps the elements with indexes i and j
func (e Entries) Swap(i, j int) {
    e[i], e[j] = e[j], e[i]
}

func WordCount(fileName string) (Entries, error) {
    // A slice of type Entry
    var entries Entries

    // Open, read and put contents of the file in a string variable
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        return entries, err
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

    // Iterate over each word
    for _, word := range words {

        // Iterate over all the words we have seen already
        exists := false
        for _,e := range entries {
            if e.Word == word {
                e.Count++
                exists = true
                break
            }
        }

        // Add a new entry in our list if this is a new word
        if !exists {
            entries = append(entries, &Entry{Word: word, Count: 1})
        }
    }

    sort.Sort(entries)
    return entries, nil
}
