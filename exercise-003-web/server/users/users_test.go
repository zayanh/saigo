package users

import (
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Create the file and append names to it
	f, err := os.OpenFile("userList_test.dat", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	f.WriteString("Kevin\n")
	f.WriteString("Lorelei\n")

	// Run setup with provided html and file just created
	Setup("home.html", "userList_test.dat")
}

func TestServer(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)
	w := httptest.NewRecorder()

	// Make sure the users were loaded correctly
	Home(w, req)
	assert.Contains(w.Body.String(), "Welcome")
	assert.Contains(w.Body.String(), "Kevin")
	assert.Contains(w.Body.String(), "Lorelei")

	// Add a new user
	req.ParseForm()
	req.Form.Set("user", "George")
	Signup(w, req)
	Home(w, req)
	assert.Contains(w.Body.String(), "George")

	// Make sure there are only 3 users
	assert.NotContains(w.Body.String(), "4.")

	// Duplicate username should not be added
	req.Form.Set("user", "George")
	Signup(w, req)
	Home(w, req)
	assert.NotContains(w.Body.String(), "4.")

	// Empty username should not be added
	req.Form.Set("user", "")
	Signup(w, req)
	Home(w, req)
	assert.NotContains(w.Body.String(), "4.")

	// Check if the redirect is working??
	assert.True(assert.HTTPRedirect(Signup, "GET", "localhost:8080/home", nil))

	// Cleanup
	os.Remove("userList_test.dat")
}
