package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

const (
	clientID     = "fakeid"
	clientSecret = "fakeSecret"
	clientURI    = "ios://fake-uri"
)

//helpers
func MockClient() *Client {
	return &Client{ClientID: clientID, ClientSecret: clientSecret, ClientURI: clientURI, Scope: "client"}

}

// Hook up gocheck into the "go test" runner.

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestNewClient(c *C) {
	client := MockClient()
	c.Assert(*NewClient(clientID, clientSecret, clientURI), Equals, *client)
}

func (s *MySuite) TestAddScope(c *C) {
	client := MockClient()
	client.AddScope("test")
	c.Assert(client.Scope, Equals, "client test")
	// scopes with whitespace are invalid!
	client.AddScope("with whitespace")
	c.Assert(client.Scope, Equals, "client test")
}

func (s *MySuite) TestAddDuplicatedScope(c *C) {
	client := MockClient()
	client.AddScope("client")
	c.Assert(client.Scope, Equals, "client")
}

func (s *MySuite) TestAddScopes(c *C) {
	client := MockClient()
	client.AddScopes("test1", "test2")
	c.Assert(client.Scope, Equals, "client test1 test2")
}

func (s *MySuite) TestSaveClient(c *C) {
	client := MockClient()
	client.SaveClient()
	newClient, _ := FetchClient(client.ClientID)
	c.Assert(*client, Equals, *newClient)
}
