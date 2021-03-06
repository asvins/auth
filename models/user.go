package models

import (
	"strconv"
	"strings"

	"github.com/asvins/common_db/redis"

	"golang.org/x/crypto/bcrypt"
)

var IDCounter int

/* CONSTRUCTION */

// User represents a device through which the user
type User struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Scope          string `json:"scope"`
	HashedPassword []byte
}

// NewUser is a constructor for users given its attributes
func NewUser(firstName, lastName, email, password, scope string) (*User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err

	}
	u := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Scope:     scope, //default scope
	}
	u.HashedPassword = hashedPass
	return u, nil
}

// SetUserID sets user's ID
func (u *User) SetFirstName(fname string) {
	u.FirstName = fname
}

// SetUserSecret sets user's secret
func (u *User) SetLastName(lname string) {
	u.LastName = lname
}

// SetUserURI sets uri for a user
func (u *User) SetEmail(email string) {
	u.Email = email
}

// AddScope adds a new scope (maybe more) to the user scopes
func (u *User) AddScope(scope string) {
	if strings.Contains(scope, " ") || strings.Contains(u.Scope, scope) {
		return
	}
	u.Scope += " " + scope
}

// AddScopes adds one or more scopes
func (u *User) AddScopes(scopes ...string) {
	for _, s := range scopes {
		u.AddScope(s)
	}
}

/* DATABASE */
func GetNextID() int {
	db := redis.NewRedisClient()
	id, err := strconv.Atoi(db.Get("id_counter").Val())
	if err != nil {
		id = 0
	}
	id += 1
	db.Set("id_counter", strconv.Itoa(id), 0)
	return id
}

// SaveUser stores user in database
func (u *User) SaveUser() error {
	db := redis.NewRedisClient()
	u.ID = GetNextID()
	return db.StoreStruct(u.Email, u)
}

// FetchUser tries to fetch an user based on an ID
func FetchUser(email string) (*User, error) {
	db := redis.NewRedisClient()
	u := User{}
	err := db.GetStruct(email, &u)
	return &u, err
}

/* LOGIC */

func AuthenticatePassword(password string, user *User) bool {
	return bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password)) == nil
}
