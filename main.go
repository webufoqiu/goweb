package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func qiu(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Qiu Sir!")
	io.WriteString(w, "<h1>hello Qiu Sir for GOlang~~</h1")
	reLaunch()
}

func main() {
	// fmt.Println("Hello World!")
	http.HandleFunc("/", qiu)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}

}
