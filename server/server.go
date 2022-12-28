package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":4000", nil)
}

func home(reponse http.ResponseWriter, request *http.Request) {
	http.ServeFile(reponse, request, "../utils/index.html")
}

