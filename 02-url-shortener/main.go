package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	urlshort "github.com/victorcete/gophercises/02-url-shortener/handler"
)

func main() {
	yamlConfig := flag.String("file", "config.yml", "YAML file with path and URL parameters")
	flag.Parse()

	yamlBytes, err := ioutil.ReadFile(*yamlConfig)
	if err != nil {
		log.Fatalln(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the MapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlBytes, mapHandler)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)
	return mux
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
