package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var ok bool

var (
	Term os.Signal = syscall.SIGTERM
)

// log the Signal to STDOUT
func logSignal(c chan os.Signal) {
	for s := range c {
		log.Println("Got signal:", s)
		ok = false
	}

}

func main() {
	ok = true
	c := make(chan os.Signal)
	signal.Notify(c, Term)

	go logSignal(c)
	log.Println("starting server at :8080")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if ok {
			log.Printf("/healthz - OK - %d", http.StatusOK)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		} else {
			log.Printf("/healthz - FAIL - %d", http.StatusServiceUnavailable)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("FAIL"))
			return
		}

	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if ok {
			log.Printf("/ready - OK - %d", http.StatusOK)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		} else {
			log.Printf("/ready - FAIL - %d", http.StatusServiceUnavailable)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("FAIL"))
			return
		}

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/- OK - %d", http.StatusOK)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome!\n"))
	})
	http.ListenAndServe(":8080", nil)
}
