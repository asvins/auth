package main

import (
	"encoding/json"
	"net/http"
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

	tkStr, err := tk.SignedString(private_key)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	body, _ := json.Marshal(&loginResponseBody{Token: tkStr, TokenType: "bearer"})
	w.Write(body)
}

type userinfoResponse struct {
	Subject  string `json:"sub"`
	Issuer   string `json:"iss"`
	IssuedAt string `json:"iat"`
	Scope    string `json:"scope"`
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

	body, _ := json.Marshal(&userinfoResponse{Subject: tk.Claims["sub"].(string), Issuer: tk.Claims["iss"].(string), IssuedAt: tk.Claims["iat"].(string), Scope: tk.Claims["scope"].(string)})
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

	usr, err := NewUser(reg.FirstName, reg.LastName, reg.Email, reg.Password, reg.Scope)
	if err != nil {
		http.Error(w, "Invalid password", 400)
		return
	}

	if usr.SaveUser() != nil {
		http.Error(w, "Error registering user. Please try again", 500)
	}

	w.Write([]byte("Success"))
}
