package main

import (
	"strings"

	"github.com/asvins/common_db"
)

// Client represents a device through which the user
type Client struct {
	ClientURI    string `json:"uri"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

// NewClient is a constructor for clients given its attributes
func NewClient(clientID, clientSecret, clientURI string) *Client {
	return &Client{
		ClientURI:    clientURI,
		ClientSecret: clientSecret,
		ClientID:     clientID,
		Scope:        "client", //default scope
	}
}

// SetClientID sets client's ID
func (c *Client) SetClientID(clientID string) {
	c.ClientID = clientID
}

// SetClientSecret sets client's secret
func (c *Client) SetClientSecret(clientSecret string) {
	c.ClientSecret = clientSecret
}

// SetClientURI sets uri for a client
func (c *Client) SetClientURI(clientURI string) {
	c.ClientURI = clientURI
}

// AddScope adds a new scope (maybe more) to the client scopes
func (c *Client) AddScope(scope string) {
	if strings.Contains(scope, " ") || strings.Contains(c.Scope, scope) {
		return
	}
	c.Scope += " " + scope
}

// AddScopes adds one or more scopes
func (c *Client) AddScopes(scopes ...string) {
	for _, s := range scopes {
		c.AddScope(s)
	}
}

// SaveClient stores Client in redis
func (c *Client) SaveClient() error {
	db := commonDB.NewRedisClient()
	return db.StoreStruct(c.ClientID, c)
}

// FetchClient tries to fetch a client based on a clientID
func FetchClient(clientID string) (*Client, error) {
	db := commonDB.NewRedisClient()
	c := Client{}
	err := db.GetStruct(clientID, &c)
	return &c, err
}
