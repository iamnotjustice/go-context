package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Operation started")
	defer log.Println("Operation ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hey, sorry it took so long")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
