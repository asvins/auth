package main

import . "gopkg.in/check.v1"

//func Test(t *testing.T) { TestingT(t) }

type CtrlSuite struct{}

var _ = Suite(&CtrlSuite{})

func CreateMockUser() {
	u, _ := NewUser(firstName, lastName, email, password, scope)
	u.SaveUser()
}

func (s *CtrlSuite) TestValidLogin(c *C) {
	// valid credentials should get a valid Token
	CreateMockUser()
	_, err := Login(email, password)
	c.Assert(err, IsNil)
}

func (s *CtrlSuite) TestInvalidLogin(c *C) {
	//invalid credentials should get an error
	_, err := Login(email, "wrongpassword")
	c.Assert(err, NotNil)
}

func (s *CtrlSuite) TestIsAuthenticatedValidToken(c *C) {
	//checks whether token is valid
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDMyNzQyODUsImlzcyI6Ind3dy5hc3ZpbnMuY29tLmJyIiwic2NvcGUiOiJwYXRpZW50Iiwic3ViIjoiam9obmRvZUBleGFtcGxlLmNvbSJ9.XYbHd0sNHnzkRlZwiTwSIdhedYKE5eKsnuJJ-D1Tnv8"

	_, err := IsAuthenticated(validToken)
	c.Assert(err, IsNil)
}

func (s *CtrlSuite) TestIsAuthenticatedInvalidToken(c *C) {
	//checks whether token is valid
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDMyNzQyODUsImlzcyI6Ind3dy5hc3ZpbnMuY29tLmJyIiwic2NvcGUiOiJwYXRpZW50Iiwic3ViIjoiam9obmRvZUBleGFtcGxlLmNvbSJ9.XYbHd0sNHnzkRlZwiTwSIdhedYKE5fKsnuJJ-D1Tnv8"

	_, err := IsAuthenticated(validToken)
	c.Assert(err, NotNil)
}
