package users

import (
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	Setup("../../")
	UserNames = make(map[string]bool)
  UserCount = 0
  f, err := os.OpenFile("userList_test.dat", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
  if err != nil {
  }
  UserFile = f
  defer UserFile.Close()
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

	// // assert.Equal(len(UserNames), 0)
	LoadUsers()
	// assert.Equal(true, UserNames["testUser"] )
	// assert.Equal(len(UserNames), 1)

	assert.Equal(2, AddUserName(""))
	assert.Equal(0, AddUserName("boom"))
	assert.Equal(1, AddUserName("boom"))
	// assert.Equal(len(UserNames), 2)

}
