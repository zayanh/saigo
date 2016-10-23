package users

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "html/template"
  "net/http"
  "bufio"
)

/* Keeping these variables global so Home and Signup can access them without them
being redefined each time. */

/*UserNames is a map that holds the list of registered users.
Was initially thinking of making userNames a slice of strings, but searching it for
duplicate entries would take O(n) time, while searching a map only takes O(1) time. */
var UserNames map[string]bool

// UserCount holds the total number of registered users
var UserCount int


// UserFile is the file handle for the list of users
var UserFile *os.File // Is this okay?


var homeT *template.Template

// Setup initializes the homeT variable which holds the HTML Template
func Setup(dir string) {
  // homeT holds the HTML template to use when loading the website
  homeT = template.Must(template.ParseFiles(dir + "/server/users/home.html"))
}

// AddUserName is called when a new user is created from the website
func AddUserName(username string) (int) {
  // Remove whitespace and check for an empty string
  if len(strings.TrimSpace(username)) <= 0 {
    return 2
  }
  // Check if the username already exists
  if _,ok := UserNames[username]; ok {
    return 1
  }

  // If above checks pass, add username to database
  UserNames[username] = true
  UserCount++
  UserFile.WriteString(username + "\n")
  return 0
}

// LoadUsers reads the provided file and builds the slice userNames
func LoadUsers() {
  // NewScanner returns a variable of type Scanner. A Scanner is a conveninent struct with
  // useful methods for reading from an io.Reader object.
  // The split function defaults to ScanLines <-- What else can I use other than bufio.ScanLines
  //    ScanBytes
  //    ScanRunes
  //    ScanWords
  scanner := bufio.NewScanner(UserFile)

  // Scan is a method for type Scanner. It generates the next "token" from scanner which can
  // be read using the Text or Bytes methods. Scan returns false when it reaches the end of
  // the input or encounters an error
  for scanner.Scan() {
    UserNames[scanner.Text()] = true
    UserCount++
  }
  // Err returns the first non-EOF error encoutnered by the scanner
  if err := scanner.Err(); err != nil {
    // handle error
  }
}

// Signup is recalled every time the user clicks on the submit button
// There is a separate function for the first time - Home
func Signup(w http.ResponseWriter, r *http.Request) {
  homeT.Execute(w, nil)
  r.ParseForm()

  username := r.Form.Get("user")
  add := AddUserName(username)
  switch add {
  case 1:
    fmt.Fprintln(w, "Username already exists!")
  case 2:
    if UserCount > 0 {
      fmt.Fprintln(w, "Username can't be empty!")
    }
  case 0:
    fmt.Fprintln(w, "User " + username + " added successfully")
  }

  // Print out the total number of and a list of all the current userNames to the screen
  // This is a dictionary so there is no inherent order...
  fmt.Fprintln(w, "There are " + strconv.Itoa(UserCount) + " current userNames:")
  i := 0 // couldn't figure out how to build the iterator into the for loop
  for val := range UserNames {
    if i > 0 {
      fmt.Fprint(w, ", ")
    }
    fmt.Fprintln(w, val)
    i++
  }
}

// Home is where the website will redirect for the very first time
func Home(w http.ResponseWriter, r *http.Request) {
  homeT.Execute(w, nil)
  r.ParseForm()

  // Print out the total number of and a list of all the current userNames to the screen
  // This is a dictionary so there is no inherent order...
  fmt.Fprintln(w, "There are " + strconv.Itoa(UserCount) + " current userNames:")
  i := 0 // couldn't figure out how to build the iterator into the for loop
  for val := range UserNames {
    if i > 0 {
      fmt.Fprint(w, ", ")
    }
    fmt.Fprintln(w, val)
    i++
  }
}
