package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	//http.ListenAndServe("", nil)
	http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)
	//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
	// open ssl
}
