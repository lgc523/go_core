package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("new watcher err:%s", err.Error())
	}
	defer w.Close()

	done := make(chan bool)
	go func() {
		for {
			defer close(done)
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err)
			}
		}
	}()
	err = w.Add("./")
	if err != nil {
		log.Fatal("add failed:", err)
	}
	<-done
}
