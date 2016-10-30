package users

import (
    // "fmt"
    "os"
    "strings"
    "html/template"
    "net/http"
    "bufio"
    "sync"
)


// View will hold the values to be passed into the HTML template
type view struct {
    Usernames map[string]bool
}

// All Usernames
var usernames map[string]bool
var usernamesMtx = &sync.Mutex{}

// userFile is the file handle for the list of users
var userFile *os.File // Is this okay?

// homeT will store the html template to be used in this website
var homeT *template.Template

// Setup initializes the homeT variable which holds the HTML Template
func Setup(dir string, filename string) {

    // homeT holds the HTML template to use when loading the website
    homeT = template.Must(template.ParseFiles(dir))

    // Initialize map of usernames
    usernames = make(map[string]bool)

    // Open file to read and write. Create if it doesn't exist. Append if it does
    f, err := os.OpenFile(filename, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }

    // use global variable
    userFile = f

    // Build the list
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
    usernamesMtx.Lock()
    defer usernamesMtx.Unlock() // is this okay?
    if _,ok := usernames[username]; ok {
        return false, true
    }

    // If above checks pass, add username to database
    usernames[username] = true
    userFile.WriteString(username + "\n")
    return false, false
}

// loadUsers reads the provided file and builds the slice UserNames
func loadUsers() {
    scanner := bufio.NewScanner(userFile)
    for scanner.Scan() {
        usernames[scanner.Text()] = true
    }

    // Err returns the first non-EOF error encoutnered by the scanner
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func copyOfUsernames() map[string]bool {
    result := make(map[string]bool)

    usernamesMtx.Lock()
    for name := range usernames {
        result[name] = true
    }
    usernamesMtx.Unlock()

    return result
}

// Signup is recalled every time the user clicks on the submit button
func Signup(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.Form.Get("user")
    empty, dup := addUserName(username)

    // Leaving this logic in future implementation of printing a message
    switch {
    case dup:
    case empty:
    default:
    }

    http.Redirect(w, r, "/home", http.StatusFound)
}


// Home is where the website will redirect everytime it wants to display all the current usernames
func Home(w http.ResponseWriter, r *http.Request) {

    // Create a View
    v := view{}
    v.Usernames = copyOfUsernames()

    // Execute the template
    homeT.Execute(w, &v)
}
