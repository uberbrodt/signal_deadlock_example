package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/uberbrodt/signal_deadlock_example/Godeps/_workspace/src/github.com/uberbrodt/imagick/imagick"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Initializing Imagemagick")
	imagick.Initialize()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit the /test route")
		imagick.Terminate()
	})

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigint, syscall.SIGTERM)

	go func() {
		for {
			<-sigint
			log.Println("Caught SIGINT, shutting down imagick environment")
			imagick.Terminate()
			os.Exit(0)
		}

	}()
	http.ListenAndServe(":4000", nil)
}
