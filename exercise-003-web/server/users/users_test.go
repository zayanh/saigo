package users

import (
	"os"
	// "fmt"
	// "strings"
	//"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Create the file and append names to it
	f, err := os.OpenFile("userList_test.dat", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		// do something
	}
	f.WriteString("Kevin\n")
	f.WriteString("Lorelei\n")
	Setup("home.html", "userList_test.dat")
}

func TestServer(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)
	w := httptest.NewRecorder()

	Home(w, req)
	assert.Contains(w.Body.String(), "Welcome")
	assert.Contains(w.Body.String(), "Kevin")
	assert.Contains(w.Body.String(), "Lorelei")

	req.ParseForm()
	req.Form.Set("user", "George")
	Signup(w, req)
	Home(w, req)
	assert.Contains(w.Body.String(), "User George added successfully")

	req.Form.Set("user", "George")
	Signup(w, req)
	Home(w, req)
	assert.Contains(w.Body.String(), "Username already exists")

	// Can't handle the string "can't"
	req.Form.Set("user", "")
	Signup(w, req)
	Home(w, req)
	assert.Contains(w.Body.String(), "be empty")

	// Check if the redirect is working??
	assert.True(assert.HTTPRedirect(Signup, "GET", "localhost:8080/home", nil))

	os.Remove("userList_test.dat")
}
