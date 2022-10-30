package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎访问 \n[URL.Path = %q]", r.URL.Path)
	Lissajous(w)
}
