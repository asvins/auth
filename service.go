package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/asvins/auth/models"
)

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponseBody struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	var l loginParams

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&l); err != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	tk, err := Login(l.Email, l.Password)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	tkStr, err := tk.SignedString(privateKey())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	body, _ := json.Marshal(&loginResponseBody{Token: tkStr, TokenType: "bearer"})
	w.Write(body)
}

type userinfoResponse struct {
	Subject  string  `json:"sub"`
	Issuer   string  `json:"iss"`
	IssuedAt float64 `json:"iat"`
	Scope    string  `json:"scope"`
}

func UserinfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	tkStr := r.Header.Get("Authorization")

	if len(tkStr) <= 0 {
		http.Error(w, "Missing token", 401)
		return
	}

	tk, err := IsAuthenticated(tkStr)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	body, _ := json.Marshal(&userinfoResponse{Subject: tk.Claims["sub"].(string), Issuer: tk.Claims["iss"].(string), IssuedAt: tk.Claims["iat"].(float64), Scope: tk.Claims["scope"].(string)})
	w.Write(body)
}

type registrationParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Scope     string `json:"scope"`
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	tkStr := r.Header.Get("Authorization")

	if len(tkStr) <= 0 {
		http.Error(w, "Missing token", 401)
		return
	}

	_, err := IsScopeAuthenticated(tkStr, "admin")
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	var reg registrationParams

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reg); err != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	usr, err := models.NewUser(reg.FirstName, reg.LastName, reg.Email, reg.Password, reg.Scope)
	if err != nil {
		http.Error(w, "Invalid password", 400)
		return
	}

	if usr.SaveUser() != nil {
		http.Error(w, "Error registering user. Please try again", 500)
		return
	}

	// Remove password for security reasons...
	usr.HashedPassword = nil
	fireEvent(EVENT_CREATED, usr)

	w.Write([]byte("Success"))
}

type discoveryResponse struct {
	Login        string `json:"login"`
	Userinfo     string `json:"userinfo"`
	Registration string `json:"registration"`
}

func DiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	cfg := LoadConfig()
	prefix := strings.Join([]string{cfg.Server.Addr, cfg.Server.Port}, ":")
	body, _ := json.Marshal(&discoveryResponse{Login: strings.Join([]string{prefix, "/api/login"}, ""), Userinfo: strings.Join([]string{prefix, "/api/userinfo"}, ""), Registration: strings.Join([]string{prefix, "/api/registration"}, "")})
	w.Write(body)
}
