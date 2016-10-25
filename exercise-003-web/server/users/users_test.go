package users

import (
	// "os"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	Setup("home.html", "userList_test.dat")
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	Home(w, req)
	assert.Contains(w.Body.String(), "Welcome")

	Signup(w, req)
	assert.Contains(w.Body.String(), "Welcome")

}
