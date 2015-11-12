package main

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	validToken   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDMyNzQyODUsImlzcyI6Ind3dy5hc3ZpbnMuY29tLmJyIiwic2NvcGUiOiJwYXRpZW50Iiwic3ViIjoiam9obmRvZUBleGFtcGxlLmNvbSJ9.XYbHd0sNHnzkRlZwiTwSIdhedYKE5eKsnuJJ-D1Tnv8"
	invalidToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDMyNzQyODUsImlzcyI6Ind3dy5hc3ZpbnMuY29tLmJyIiwic2NvcGUiOiJwYXRpZW50Iiwic3ViIjoiam9obmRvZUBleGFtcGxlLmNvbSJ9.XYbHd0sNHnzkRlZwiTwSIdhedYKE5fKsnuJJ-D1Tnv8"
)

type mockResource struct {
	owner string
}

// Implementing Resource
func (r mockResource) Owner() string {
	return r.owner
}

func TestIsAuthorized(t *testing.T) {
	authorizePatient := func(tk jwt.Token) bool {
		return tk.Claims["scope"] == "patient"
	}

	Convey("When authorizing resources", t, func() {

		Convey("An invalid token will return false", func() {
			So(IsAuthorized(invalidToken, authorizePatient), ShouldEqual, false)
		})

		Convey("A valid token with a true predicate should return true", func() {
			So(IsAuthorized(validToken, authorizePatient), ShouldEqual, true)
		})
	})
}

func TestIsOwner(t *testing.T) {
	r1 := mockResource{owner: "johndoe@example.com"}
	r2 := mockResource{owner: "dilmita@example.com"}
	Convey("When checking resource ownership", t, func() {

		Convey("We recognize the resource owner", func() {
			So(IsOwner(validToken, r1), ShouldEqual, true)
		})

		Convey("Everybody else should not be recognized", func() {
			So(IsOwner(validToken, r2), ShouldEqual, false)
		})
	})
}

func TestHasScope(t *testing.T) {
	Convey("When validating scopes", t, func() {
		Convey("We return true to a matching scope", func() {
			So(HasScope(validToken, "admin"), ShouldEqual, false)
			So(HasOneScope(validToken, []string{"admin", "pharmacist"}), ShouldEqual, false)
		})

		Convey("We deny a non matching scope", func() {
			So(HasScope(validToken, "patient"), ShouldEqual, true)
			So(HasOneScope(validToken, []string{"admin", "patient"}), ShouldEqual, true)
		})
	})
}
