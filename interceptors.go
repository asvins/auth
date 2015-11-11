package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Resource interface {
	Owner() string
}

func IsAuthorized(token string, pred func(t jwt.Token) bool) bool {
	key := LoadConfig().Service.Private_Key
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if ok := token.Method; ok != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	return err == nil && pred(*tk)
}

func IsOwner(token string, r Resource) bool {
	return IsAuthorized(token, func(t jwt.Token) bool {
		return t.Claims["sub"].(string) == r.Owner()
	})
}

func HasOneScope(token string, scopes []string) bool {
	return IsAuthorized(token, func(t jwt.Token) bool {
		scope := t.Claims["scope"].(string)
		for _, s := range scopes {
			if strings.Contains(scope, s) {
				return true
			}
		}
		return false
	})
}

func HasScope(token, scope string) bool {
	return IsAuthorized(token, func(t jwt.Token) bool {
		return strings.Contains(t.Claims["scope"].(string), scope)
	})
}
