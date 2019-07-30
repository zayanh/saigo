package count

import (
	"strings"
	"io/ioutil"
	"sort"
	"strconv"
)

// Each line in the final output will contain a word
//   and the word's count
type Line struct {
	Word string
	Count int
}
type Lines []Line

// Method to search the Lines for an existing word
//    and return the index if found or -1 if not
func (l Lines) Find(value string) int {
	for i, v := range l {
		if v.Word == value {
			return i
		}
	}
	return -1
}


// Go's sorting setup
// ByCount implements sort.Interface for []Line based on
//    the Count field
type ByCount Lines

// Teach sort how to measure the length of our structure
func (a ByCount) Len() int				{ return len(a) }

// Teach sort how to swap two elements in our structure
func (a ByCount) Swap(i, j int)			{ a[i], a[j] = a[j], a[i] }
// Teach sort what order to sort in
func (a ByCount) Less(i, j int) bool	{ return a[i].Count > a[j].Count }


// Create a custom set of delimiters to split our file on
//    can we use regexp? (\s)
func Split(r rune) bool {
	return r == ' ' || r == '\n'
}



func Count(filename string) string {
	// object to hold final answer
	var lines Lines

	// read the file and store it as a string
	bs, err := ioutil.ReadFile(filename) //("7oldsamr.txt")
	if err != nil {
		return ""
	}
	file := string(bs)

	// remove punctuation from file
	//  can we use regexp? (\W)
	file = strings.ReplaceAll(file, "\t", "")
	file = strings.ReplaceAll(file, ",", "")
	file = strings.ReplaceAll(file, ".", "")
	file = strings.ReplaceAll(file, "\"", "")
	file = strings.ReplaceAll(file, ":", "")
	file = strings.ReplaceAll(file, "?", "")
	file = strings.ReplaceAll(file, "!", "")
	file = strings.ReplaceAll(file, "    ", "")

	// make everything lower case
	file = strings.ToLower(file)

	// split across every space and newline
	words := strings.FieldsFunc(file, Split)

	// count the words
	for _,val := range words {

		// check if we've already come across this word
		//    if yes, increment count
		//    if no, add word to our slice
		index := lines.Find(val)
		if index >= 0 {
			lines[index].Count++
		} else {
			lines = append(lines, Line{val, 1})
		}
	}

	// sort our structure based on how we taught it earlier
	sort.Sort(ByCount(lines))

	outString := ""
	for _, val := range lines {
		outString = outString + strconv.Itoa(val.Count) + " " + val.Word + "\r\n"
	}

	return outString
}
