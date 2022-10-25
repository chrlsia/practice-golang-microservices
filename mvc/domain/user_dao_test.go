package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Errorf("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Errorf("we were expecting an error when user id is 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Errorf("we were expectin 404 when user was not found")
	}
}

/*
go test
go test -cover
*/
