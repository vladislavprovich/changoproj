package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func requestone(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "request_one\n")
}

func requesttwo(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Fprintf(w, "POST request\n")
	case "GET":
		fmt.Fprintf(w, "GET request\n")
	case "PUT":
		fmt.Fprintf(w, "PUT request\n")
	case "DELETE":
		fmt.Fprintf(w, "DELETE request\n")

	}
}

func requesthre(w http.ResponseWriter, req *http.Request) {
	url := "https://github.com/vladislavprovich"

	res, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	html := string(body)
	if strings.Contains(html, "Popular repositories") {
		fmt.Fprintf(w, "Popular repositories")
	} else {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	}
}

func getyoutube(w http.ResponseWriter, req *http.Request) {
	url := "https://www.youtube.com"

	res, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	fmt.Fprintf(w, string(body))
	//html := string(body)

}

func main() {

	http.HandleFunc("/requestone", requestone)
	http.HandleFunc("/requesttwo", requesttwo)
	http.HandleFunc("/checkgithub", requesthre)
	http.HandleFunc("/getyoutube", getyoutube)

	http.ListenAndServe(":8080", nil)
}
