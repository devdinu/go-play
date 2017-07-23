package main

import "net/http"

func main() {
	http.HandleFunc("/ping", Authenticator(Ping))
	http.ListenAndServe(":8080", nil)

	testGopherThing()
}
