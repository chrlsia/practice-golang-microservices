package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

	//if user != nil {
	//	t.Errorf("we were not expecting a user with id 0")
	//}

	//if err == nil {
	//	t.Errorf("we were expecting an error when user id is 0")
	//}

	//if err.StatusCode != http.StatusNotFound {
	//	t.Errorf("we were expectin 404 when user was not found")
	//}
}

/*
go test
go test -cover

------
go get -u github.com/stretchr/testify/assert
import "github.com/stretchr/testify/assert"
*/
