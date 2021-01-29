package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if val := q.Get("delay"); val != "" {
		t, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("Error: %v", err)
			fmt.Fprintf(w, "Failed to convert delay value")
		}
		log.Printf("Sleep: %d seconds", t)
		time.Sleep(time.Duration(t) * time.Second)
	}

	fmt.Fprintf(w, "Hello World! %s UTC\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
}

func bug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "crashed!")
	log.Fatalf("crashed")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/bug", bug)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Failed to listen on port:%v", err)
	}
}
