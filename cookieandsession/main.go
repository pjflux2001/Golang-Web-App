package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var userStore = map[string]string{"Wallace": "badPassword"}
var sessionStore = map[string]string{} // preferably uuid as key but username would work here
var sessionTtl = 5                     // in seconds
var fileName = "login.html"
var secret = "secret"

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err == nil {
		_, ok := sessionStore[cookie.Value]
		if ok {
			fmt.Fprintf(w, "You've already logged in.")
			return
		}
	}

	display(w, "Please login")
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if userStore[username] != password {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Credential not found")
		return
	}

	cookie, err := r.Cookie("session-id")
	if err != nil {
		cookie = &http.Cookie{
			Name: "session-id",
		}
	}
	cookie.Value = username
	code := getCode(cookie.Value)
	sessionStore[code] = username

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've logged in. Welcome to Golang Dojo")
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err == nil {
		code := getCode(cookie.Value)
		_, ok := sessionStore[code]
		if ok {
			delete(sessionStore, code)
			fmt.Fprintf(w, "You've successfully logged out.")
			return
		}
	}
	fmt.Fprintf(w, "You never logged in.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	case "/logout":
		logout(w, r)
	case "/home":
		fmt.Fprintf(w, "Welcome to Golang Dojo")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not implemented")
	}
}

func display(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
	//http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)
	// leave a comment if you want a video on https
	// even though it's not technically quite directly related to Go
	//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
}
