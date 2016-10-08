package main

import (
	"fmt"
	"github.com/axet/desktop/go"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[1:]
	log.Println(file)
	if file == "" {
		file = "./index.html"
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/", handler)
	go func() {
		desktop.BrowserOpenURI("http://localhost:8080")
	}()
	http.ListenAndServe(":8080", nil)
}
