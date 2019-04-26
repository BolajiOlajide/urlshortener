package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	urlshort "github.com/BolajiOlajide/urlshortener"
)

func main() {
	yamlFileName := flag.String("yml", "routes.yml", "Specify a yaml file to use")
	jsonFileName := flag.String("json", "routes.json", "specify JSON to load route from")
	mux := defaultMux()
	flag.Parse()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile(*yamlFileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	json, err := ioutil.ReadFile(*jsonFileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
