package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/asvins/router"
	"github.com/asvins/router/errors"
	"github.com/dgrijalva/jwt-go"
)

type Resource interface {
	Owner() string
}

type isAuthorizedInterceptor func(t jwt.Token) bool

func (i isAuthorizedInterceptor) Intercept(w http.ResponseWriter, r *http.Request) errors.Http {
	tokenString := r.Header.Get("Authorization")
	if !IsAuthorized(tokenString, i) {
		return errors.Unauthorized("Unauthorized")
	}
	return nil
}

func IsAuthorizedInterceptor(f func(t jwt.Token) bool) router.Interceptor {
	return isAuthorizedInterceptor(f)
}

// type isOwnerInterceptor  Resource
//
// func (i isOwnerInterceptor) Intercept(w http.ResponseWriter, r *http.Request) errors.Http {
// 	tokenString := r.Header.Get("Authorization")
// 	if !IsOwner(tokenString, i) {
// 		return errors.Unauthorized("Unauthorized")
// 	}
// 	return nil
// }
//
// func IsOwnerInterceptor(r Resource) router.Interceptor {
// 	return isOwnerInterceptor(r)
// }

type hasScopeInterceptor string

func (i hasScopeInterceptor) Intercept(w http.ResponseWriter, r *http.Request) errors.Http {
	tokenString := r.Header.Get("Authorization")
	if !HasScope(tokenString, string(i)) {
		return errors.Unauthorized("Unauthorized")
	}
	return nil
}

func HasScopeInterceptor(scope string) router.Interceptor {
	return hasScopeInterceptor(scope)
}

type hasOneScopeInterceptor []string

func (i hasOneScopeInterceptor) Intercept(w http.ResponseWriter, r *http.Request) errors.Http {
	tokenString := r.Header.Get("Authorization")
	if !HasOneScope(tokenString, []string(i)) {
		return errors.Unauthorized("Unauthorized")
	}
	return nil
}

func HasOneScopeInterceptor(scopes []string) router.Interceptor {
	return hasOneScopeInterceptor(scopes)
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
