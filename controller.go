package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// q horror
var private_key = []byte("contumacia")

const Issuer = "www.asvins.com.br"

func Login(email, password string) (*jwt.Token, error) {
	user, err := AuthenticateUser(email, password)

	if err != nil {
		return nil, err
	}

	return issueToken(user.Email, user.Scope)
}

func AuthenticateUser(email, password string) (*User, error) {
	if len(password) < 8 {
		return nil, fmt.Errorf("Password too short. Please use at least 8 characters.")
	}

	user, err := FetchUser(email)
	combinationErr := fmt.Errorf("Please check your email/password combination")

	if err != nil || !AuthenticatePassword(password, user) {
		return nil, combinationErr
	}

	return user, nil
}

func IsAuthenticated(tokenStr string) error {
	_, err := validateToken(tokenStr)
	return err
}

func validateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if ok := token.Method; ok != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return private_key, nil
	})
	return token, err
}

func issueToken(email, scope string) (*jwt.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["iss"] = Issuer
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = email
	token.Claims["scope"] = scope
	return token, nil
}
