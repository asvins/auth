package main

import (
	"github.com/c0rrzin/router"
)

// DefRoutes is the function in which all routes are created
func DefRoutes() {
	router.DefRoute("GET", "/api/discovery", DiscoveryHandler)
	router.DefRoute("POST", "/api/login", LoginHandler)
	router.DefRoute("GET", "/api/userinfo", UserinfoHandler)
	router.DefRoute("POST", "/api/registration", RegistrationHandler)
}
