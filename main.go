package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

//go:embed views/*
var viewsFS embed.FS

func main() {

	tmpls, err := template.New("").
		ParseFS(viewsFS,
			"views/*.html",
		)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		envmap := make(map[string]string)
		for _, e := range os.Environ() {
			ep := strings.SplitN(e, "=", 2)
			envmap[ep[0]] = ep[1]
		}
		err = tmpls.ExecuteTemplate(w, "envlist.html", struct {
			Env map[string]string
		}{
			Env: envmap,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			log.Fatal("debug.ReadBuildInfo() failed")
		}
		w.Header().Set("X-Version", buildInfo.Main.Version)

		err = tmpls.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
