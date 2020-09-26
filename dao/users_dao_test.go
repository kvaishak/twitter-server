package dao

import (
	"net/http"
	"testing"
)

func TestWrongUserName(t *testing.T) {
	userdata, err := GetUserData("vaishak")

	if userdata != nil {
		t.Error("User data should be nil for username with value vaishak")
	}
	if err == nil {
		t.Error("Error should not be NILL for username vaishak")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("Error status code must be 404 for username: vaishak")
	}
}

func TestCorrectUserName(t *testing.T) {
	userdata, err := GetUserData("mark")

	if userdata == nil {
		t.Error("User data should not be NIL for username with value mark")
	}
	if err != nil {
		t.Error("Error should be NILL for username Mark")
	}

}
