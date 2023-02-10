package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello/", hello)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}

func handler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("404 Page Not Found")); err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler404(w, r)
		return
	}
	user := User{Id: 1, Name: "Ivan"}
	bytes, _ := json.MarshalIndent(user, "", "   ")
	if _, err := w.Write([]byte(bytes)); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	pathRegexp := regexp.MustCompile(`^/hello/\w+$`)
	if !pathRegexp.Match([]byte(r.URL.Path)) {
		handler404(w, r)
		return
	}

	name := strings.Split(r.URL.Path, "/")[2]
	if _, err := w.Write([]byte(fmt.Sprintf("Hello, %s", name))); err != nil {
		panic(err)
	}
}

type User struct {
	Id   int
	Name string
}
