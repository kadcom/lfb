package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	cfg := NewConfig()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := welcome()
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("X-Powered-By", "Sega Genesis")
		w.WriteHeader(http.StatusOK)
		component.Render(r.Context(), w)
	})

	addr := ":" + strconv.Itoa(int(cfg.Port))

	log.Println("Listening on", addr)
	http.ListenAndServe(addr, nil)
}
