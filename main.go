package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"math/rand"
)

//go:embed views/*
var viewsFS embed.FS

func main() {
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		greetings := []string{"hello", "hi", "howdy", "hola", "bonjour", "ciao", "hallo", "hej", "guten tag", "namaste", "ni hao", "salut", "merhaba", "こんにちは", "안녕하세요", "안녕", "你好", "नमस्ते", "П��ивет", "你好", "Olá", "Hallo", "Ciao", "Hola", "Bonjour", "Merhaba", "こんにちは", "안녕하세요", "안녕", "你好", "नमस्ते", "Привет", "你好", "Olá", "Hallo", "Ciao", "Hola", "Bonjour", "Merhaba", "こんにちは", "안녕하세요", "안녕", "你好", "नमस्ते", "Привет", "你好", "Olá", "Hallo", "Ciao", "Hola", "Bonjour", "Merhaba"}
		randomGreeting := greetings[rand.Intn(len(greetings))]
		w.Write([]byte("<div id=\"parent-div\"><h1>" + randomGreeting + "</h1></div>"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpls, err := template.New("").
			ParseFS(viewsFS,
				"views/*.html",
			)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpls.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
