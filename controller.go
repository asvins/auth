package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/asvins/auth/models"
	"github.com/dgrijalva/jwt-go"
)

func Login(email, password string) (*jwt.Token, error) {
	user, err := AuthenticateUser(email, password)

	if err != nil {
		return nil, err
	}

	return issueToken(user.ID, user.Email, user.Scope)
}

func AuthenticateUser(email, password string) (*models.User, error) {
	if len(password) < 8 {
		return nil, fmt.Errorf("Password too short. Please use at least 8 characters.")
	}

	user, err := models.FetchUser(email)
	combinationErr := fmt.Errorf("Please check your email/password combination")

	if err != nil || !models.AuthenticatePassword(password, user) {
		return nil, combinationErr
	}

	return user, nil
}

func IsAuthenticated(tokenStr string) (*jwt.Token, error) {
	return validateToken(tokenStr)
}

func IsScopeAuthenticated(tokenStr, scope string) (*jwt.Token, error) {
	tk, err := validateToken(tokenStr)
	if err != nil {
		return nil, err
	}
	usrScope := tk.Claims["scope"].(string)
	if !strings.Contains(usrScope, scope) {
		return nil, fmt.Errorf("Forbidden")
	}
	return tk, nil
}

func validateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if ok := token.Method; ok != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey(), nil
	})
	return token, err
}

func issueToken(id int, email, scope string) (*jwt.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["iss"] = LoadConfig().Service.Issuer
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = email
	token.Claims["scope"] = scope
	token.Claims["id"] = id
	return token, nil
}
