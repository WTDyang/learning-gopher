package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]float32

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("请求地址为：%s", r.RemoteAddr)
	if r.URL.Path != "/" && r.URL.Path != "/price" {
		msg := fmt.Sprintf("no such page: %s\n", r.URL)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}
	if r.URL.Path == "/price" {
		item := r.URL.Query().Get("item")
		if item == "" {
			fmt.Fprintf(w, "请指定商品")
			return
		}
		fmt.Fprintf(w, "%s: %.2f$\n", item, db[item])
		return
	}
	for item, price := range db {
		fmt.Fprintf(w, "%s: %f\n", item, price)
	}
}

func main() {
	db := database{"pen": 5.2, "ruler": 3.5}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
