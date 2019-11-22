package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("UTC")
	t := time.Now().In(loc).Format(time.RFC3339)
	fmt.Fprintf(w, "Please submit a screen shot of this page with your solution %s! %v", r.URL.Path[1:], t)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
