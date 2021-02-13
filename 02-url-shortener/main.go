package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	urlshort "github.com/victorcete/gophercises/02-url-shortener/handler"
)

func main() {
	mode := flag.String("mode", "yaml", "config mode, can be `json` or `yaml`")
	configFile := flag.String("file", "config.yaml", "file with path and URL parameters, can be `json` or `yaml`")
	flag.Parse()

	configBytes, err := ioutil.ReadFile(*configFile)
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

	var fileHandler http.HandlerFunc

	switch *mode {
	case "json":
		// Build the YAMLHandler using the MapHandler as the fallback
		jsonHandler, err := urlshort.JSONHandler(configBytes, mapHandler)
		if err != nil {
			log.Fatalln(err)
		}
		fileHandler = jsonHandler
	case "yaml":
		// Build the YAMLHandler using the MapHandler as the fallback
		yamlHandler, err := urlshort.YAMLHandler(configBytes, mapHandler)
		if err != nil {
			log.Fatalln(err)
		}
		fileHandler = yamlHandler
	default:
		flag.Usage()
		os.Exit(42)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", fileHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)
	return mux
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
