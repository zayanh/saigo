package WordCount

import (
    // "fmt"
    "testing"
)

func TestLess(t *testing.T) {
    var testEntries Entries
    for i := 0; i < 2; i++ {
        testEntries = append(testEntries, &Entry{Word: string(byte(0x61+i)), Count: i})
    }
    b := testEntries.Less(0,1)
    if b != false {
        t.Error("Expected true, got ", b)
    }
}

func TestSwap(t *testing.T) {
    var testEntries Entries
    for i := 0; i < 2; i++ {
        testEntries = append(testEntries, &Entry{Word: string(byte(0x61+i)), Count: i})
    }
    testEntries.Swap(0,1)
    if testEntries[0].Count == 0 {
        t.Error("Expected testEntries[0].Count to be 1, got ", testEntries[0].Count)
    }
}

// Would it be better to get the filename from the command line in the test as well?
func TestSort(t *testing.T) {
    _, testEntries := WordCount("7oldsamr.txt")

    // Is there a cleaner way to check the next value in the slice?
    // if I used "for i,val := range testEntries"
    for i := 0; i < len(testEntries)-1; i++ {
        if testEntries[i+1].Count > testEntries[i].Count {
            t.Error("Expected count for", testEntries[i+1].Word, "(",
                testEntries[i+1].Count, ") to be less than count for",
                testEntries[i].Word, "(", testEntries[i].Count, ")" )
            break
        }
    }
}

func TestUniqueWord(t *testing.T) {
    _, testEntries := WordCount("7oldsamr.txt")
    err := 0
    for i := 0; i < len(testEntries)-1; i++ {
        for j := i+1; j < len(testEntries); j++ {
            if testEntries[i].Word == testEntries[j].Word {
                t.Error("Found two separate entries with the word:", testEntries[i].Word)
                err = 1
                break
            }
        }
        if err == 1 {break}
    }
}

// What else can we benchmark in this function?
func BenchmarkWordCount(b *testing.B) {
    for n := 0; n < b.N; n++ {
        WordCount("7oldsamr.txt")
    }
}