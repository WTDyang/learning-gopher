package main

import (
	"fmt"
	"log"
	"net/http"
)

type database2 map[string]float32

func (db database2) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %f\n", item, price)
	}
	return
}

func (db database2) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if item == "" {
		fmt.Fprintf(w, "请指定商品")
		return
	}
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %.2f$\n", item, price)
	return
}
func main() {
	db := database2{"pen": 5.2, "ruler": 3.5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
