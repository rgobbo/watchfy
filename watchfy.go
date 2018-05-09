package watchfy

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type fnHandler func(string)

//NewWatcher - Create a new watcher files
func NewWatcher(pathToWatch []string, showLog bool, fn fnHandler) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		var strevent = ""
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write && strevent != time.Now().Format("Mon Jan _2 15:04:05 2006")+event.String() {
					log.Println("Modified file:", event.Name)
					fn(event.Name)
					strevent = time.Now().Format("Mon Jan _2 15:04:05 2006") + event.String()
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	currentRoot, _ := os.Getwd()
	for _, s := range pathToWatch {
		completePath := path.Join(currentRoot, s)

		inf, err := os.Stat(completePath)
		if err != nil {
			log.Println("Watcher : ", err)
		}

		if inf.IsDir() == true { // If is dir path
			info, err := ReadAllFiles(completePath)
			for i, ff := range info {
				if ff.IsDir() {
					if showLog == true {
						log.Println("Watching folder :", i)
					}

					err = watcher.Add(i)
					if err != nil {
						log.Println("Watcher adding path error:", err.Error())
					}
				}
			}

		} else {
			if showLog == true {
				log.Println("Watching file :", completePath)
			}
			err = watcher.Add(completePath)
			if err != nil {
				log.Println("Watcher adding path error:", err.Error())
			}
		}

		if err != nil {
			log.Println(err)
		}
	}

	<-done
}

//ReadAllFiles - read files from folder
func ReadAllFiles(dirname string) (map[string]os.FileInfo, error) {
	list := make(map[string]os.FileInfo, 0)

	err := filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		list[path] = f
		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}