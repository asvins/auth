package models

import (
	"testing"

	. "gopkg.in/check.v1"
)

const (
	UserID    = "fakeid"
	firstName = "John"
	lastName  = "Doe"
	email     = "johndoe@example.com"
	scope     = "patient"
	password  = "l337sp34k"
)

//helpers
func MockUser() *User {
	return &User{ID: UserID, FirstName: firstName, LastName: lastName, Scope: "user", Email: email}
}

// Hook up gocheck into the "go test" runner.

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAddScope(c *C) {
	user := MockUser()
	user.AddScope("test")
	c.Assert(user.Scope, Equals, "user test")
	// scopes with whitespace are invalid!
	user.AddScope("with whitespace")
	c.Assert(user.Scope, Equals, "user test")
}

func (s *MySuite) TestAddDuplicatedScope(c *C) {
	user := MockUser()
	user.AddScope("user")
	c.Assert(user.Scope, Equals, "user")
}

func (s *MySuite) TestAddScopes(c *C) {
	user := MockUser()
	user.AddScopes("test1", "test2")
	c.Assert(user.Scope, Equals, "user test1 test2")
}

func (s *MySuite) TestSaveUser(c *C) {
	user := MockUser()
	user.SaveUser()
	newUser, _ := FetchUser(user.Email)
	c.Assert((*user).ID, Equals, (*newUser).ID)
}
