package main

import (
	"github.com/rgobbo/watchfy"
	"path"
	"net/http"
	"log"
)

func main (){
	go watchfy.NewWatcher([]string{"sample"}, true, func(filename string) {
		ext := path.Ext(filename)
		if ext == ".html" {
			// Do something
			log.Println("The file:", filename, " was modified.")
		}
	})

	fs := http.FileServer(http.Dir("sample"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}