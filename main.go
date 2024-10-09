package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
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

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Помилка при завантаженні сторінки:", err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Помилка при парсингу HTML:", err)
	}
	doc.Find("h2.f4.mb-2.text-normal").Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		fmt.Fprintf(w, "Знайдений фрагмент HTML:")
		fmt.Fprintf(w, html)
	})

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
	html := string(body)

	if strings.Contains(html, "Назад") {
		fmt.Fprintf(w, "<h1>Назад<h1>")
	} else {
		fmt.Fprintf(w, "error")
	}

}

func main() {

	http.HandleFunc("/requestone", requestone)
	http.HandleFunc("/requesttwo", requesttwo)
	http.HandleFunc("/checkgithub", requesthre)
	http.HandleFunc("/getyoutube", getyoutube)

	http.ListenAndServe(":8080", nil)
}
