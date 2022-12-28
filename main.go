/*package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}*/

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, "Hello World from path: %s\n", request.URL.Path)
	http.ServeFile(response, request, "./utils/index.html")
}