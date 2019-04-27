package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"database/sql"

	urlshort "github.com/BolajiOlajide/urlshortener"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	yamlFileName := flag.String("yml", "routes.yml", "Specify a yaml file to use")
	jsonFileName := flag.String("json", "routes.json", "specify JSON to load route from")
	mux := defaultMux()
	flag.Parse()

	db, err := sql.Open("mysql", "bolaji:andela@/demodb")
	checkErr(err)

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile(*yamlFileName)
	checkErr(err)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	json, err := ioutil.ReadFile(*jsonFileName)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT * FROM urls")
	checkErr(err)

	dbPaths := make(map[string]string)
	for rows.Next() {
		var path, url string
		var id int
		err := rows.Scan(&id, &path, &url)
		checkErr(err)
		dbPaths[path] = url
	}

	// HANDLERS
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	checkErr(err)

	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	checkErr(err)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
