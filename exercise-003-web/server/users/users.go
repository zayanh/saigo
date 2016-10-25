package users

import (
  // "fmt"
  "os"
  "strings"
  // "strconv"
  "html/template"
  "net/http"
  "bufio"
)

/* Keeping these variables global so Home and Signup can access them without them
being redefined each time. */

type data struct {
  /*UserNames is a map that holds the list of registered users.
  Was initially thinking of making UserNames a slice of strings, but searching it for
  duplicate entries would take O(n) time, while searching a map only takes O(1) time. */
  UserNames map[string]bool

  // Message will inform the user of the results from the previous attempt to add a new user
  Message string
}

var d data

// userFile is the file handle for the list of users
var userFile *os.File // Is this okay?

// homeT will store the html template to be used in this website
var homeT *template.Template

// Setup initializes the homeT variable which holds the HTML Template
func Setup(dir string, filename string) {
  // homeT holds the HTML template to use when loading the website
  homeT = template.Must(template.ParseFiles(dir))

  // d = data{}

  // Initialize map of usernames
  d.UserNames = make(map[string]bool)

  // Open file to read and write. Create if it doesn't exist. Append if it does
  f, err := os.OpenFile(filename, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
  if err != nil {
    //handle error
  }
  userFile = f // use global variable -- is this approach okay?
  // defer userFile.Close()

  //Build the list
  loadUsers()
}

// addUserName is called when a new user is created from the website
// first bool is if username is empty, second bool is if username exists
func addUserName(username string) (bool, bool) {
  // Remove whitespace and check for an empty string
  if len(strings.TrimSpace(username)) <= 0 {
    return true, false
  }
  // Check if the username already exists
  if _,ok := d.UserNames[username]; ok {
    return false, true
  }

  // If above checks pass, add username to database
  d.UserNames[username] = true
  userFile.WriteString(username + "\n")
  return false, false
}

// loadUsers reads the provided file and builds the slice UserNames
func loadUsers() {
  // NewScanner returns a variable of type Scanner. A Scanner is a conveninent struct with
  // useful methods for reading from an io.Reader object.
  // The split function defaults to ScanLines <-- What else can I use other than bufio.ScanLines
  //    ScanBytes
  //    ScanRunes
  //    ScanWords
  scanner := bufio.NewScanner(userFile)

  // Scan is a method for type Scanner. It generates the next "token" from scanner which can
  // be read using the Text or Bytes methods. Scan returns false when it reaches the end of
  // the input or encounters an error
  for scanner.Scan() {
    d.UserNames[scanner.Text()] = true
  }
  // Err returns the first non-EOF error encoutnered by the scanner
  if err := scanner.Err(); err != nil {
    // handle error
  }
}

// Signup is recalled every time the user clicks on the submit button
func Signup(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("user")
  empty, dup := addUserName(username)
  switch {
  case dup:
    // fmt.Fprintln(w, "Username already exists!")
    d.Message = "Username already exists!"
  case empty:
    // fmt.Fprintln(w, "Username can't be empty!")
    d.Message = "Username can't be empty!"
  default:
    // fmt.Fprintln(w, "User " + username + " added successfully")
    d.Message = "User " + username + " added successfully"
  }

  // homeT.Execute(w, &UserNames)
  http.Redirect(w, r, "/home", http.StatusFound)

  // Redirect doesn't work with fmt.Fprintln .... not quite sure why...
  //    "http: multiple response.WriteHeader calls"
}

// Home is where the website will redirect for the very first time only
func Home(w http.ResponseWriter, r *http.Request) {
  homeT.Execute(w, &d)
  d.Message = ""
}
