// Exercise 3 - Build a web server that accepts new userNames and keeps track of existing ones
package main

import (
  "os"
  "net/http"

  "github.com/zayanh/saigo/exercise-003-web/server/users"
)

func main() {
  // Initialize the map of usernames
  users.UserNames = make(map[string]bool)
  users.UserCount = 0

  // Call the Setup function which initializes homeT
  users.Setup(".")

  // Open file to read and write. Create if it doesn't exist. Append if it does
  f, err := os.OpenFile("server/userList.dat", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
  if err != nil {
    //handle error
  }
  users.UserFile = f // use global variable -- is this approach okay?
  defer users.UserFile.Close()

  // build the list
  users.LoadUsers()

  // Setup the various URIs to accept
  http.HandleFunc("/signup", users.Signup)
  http.HandleFunc("/home", users.Home)
  http.ListenAndServe(":8080", nil)
}
