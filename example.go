package main

import (
	"github.com/uberbrodt/signal_deadlock_example/Godeps/_workspace/src/github.com/uberbrodt/imagick/imagick"
	"log"
	"net/http"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Initializing Imagemagick")
	imagick.Initialize()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit the /test route")
		imagick.Terminate()
	})
	http.ListenAndServe(":4000", nil)
}
