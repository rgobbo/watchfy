# watchfy

A golang library to watch file changes and call a function

Using github.com/fsnotify/fsnotify

## Installation

As a library

```shell
go get github.com/rgobbo/watchfy
```

## Usage

```go
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

```

## Documentation

### NewWatcher(pathToWatch []string, showLog bool, fn fnHandler)
Create a new watcher files

Parameters :

pathToWatch : array of strings containing paths to watch file changes

showLog : if true show log information if false log information will be ommited

fn: Function handler func(filename string) with parameter filename that will be recived from the file change event

## Sample

To run sample execute go run main.go and change something in the  index.html file and save.

You will see the log files.