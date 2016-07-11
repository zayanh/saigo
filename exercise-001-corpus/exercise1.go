package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//QUESTION: If I don't manually parse for the newline,
//  they don't appear if the line doesn't start with whitespace

type MapWithOrder struct {
    baseMap map[string]int
    orderSlice []string
}

func main() {
    //open, read and put contents of the file in a string variable
    byteString, err := ioutil.ReadFile("test")
    if err != nil {
        fmt.Println("Error reading file")
        return
    }
    str := string(byteString)

    //get rid of punctuation
    str = strings.Replace(str, "\"", "", -1)
    str = strings.Replace(str, ",", "", -1)
    str = strings.Replace(str, "!", "", -1)
    str = strings.Replace(str, ":", "", -1)
    str = strings.Replace(str, "-", "", -1)
    str = strings.Replace(str, ".", "", -1)
    str = strings.Replace(str, "\t", "", -1)

    //make it all lower case to compare
    str = strings.ToLower(str)

    //split str into a slice of type string at every instance of " "
    str2 := strings.Split(str, " ")

    //
    m := MapWithOrder{ baseMap: make(map[string]int), orderSlice: make([]string, len(str2)) }

    for _, val := range str2 {

        // watchout for newlines!
        str3 := strings.Split(val, "\n")

        for _, val2 := range str3 {

            //multiple sapces give null strings
            if val2 != "" {
                m.baseMap[val2]++

                //There are two scenarios here, the entry exists or it doesn't
                // if it's a new entry, add it to the end:
                currentIndex := len(m.orderSlice)
                // fmt.Println(currentIndex)
                m.orderSlice[currentIndex-1] = val2

                //infinite loop until order is correct
                for {
                    if (currentIndex > 2) && 
                       (m.baseMap[m.orderSlice[currentIndex-1]] > m.baseMap[m.orderSlice[currentIndex-2]]) {
                        //swap
                        tmp := m.orderSlice[currentIndex-1]
                        m.orderSlice[currentIndex-1] = m.orderSlice[currentIndex-2]
                        m.orderSlice[currentIndex-2] = tmp
                        break
                    }
                    currentIndex--
                }

                // if the entry does exist

            }
        }
    }

    for _, val := range m.orderSlice {
        fmt.Println(m.baseMap[val], val)
    }
}
