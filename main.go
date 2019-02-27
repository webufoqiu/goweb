package main

import (
	"io"
	"log"
	"net/http"
)

func qiu(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Qiu Sir!")
	io.WriteString(w, "<h1>hello Qiu Sir for GOlang~~</h1")
}

func main() {
	// fmt.Println("Hello World!")
	http.HandleFunc("/", qiu)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}

}
