// Exercise 3 - Build a web server that accepts new userNames and keeps track of existing ones
package main

import (
  "net/http"

  "github.com/zayanh/saigo/exercise-003-web/server/users"
)

func main() {
  // Call the Setup function which initializes homeT and initializes the map of usernames
  users.Setup("./server/users/home.html", "./server/userList.dat")

  // Setup the various URIs to accept
  http.HandleFunc("/signup", users.Signup)
  http.HandleFunc("/home", users.Home)
  http.ListenAndServe(":8080", nil)
}
