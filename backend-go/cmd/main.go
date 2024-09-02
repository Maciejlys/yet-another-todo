package main

import (
	"log"
	"net/http"

	"github.com/Maciejlys/yet-another-todo/db"
	"github.com/Maciejlys/yet-another-todo/web"
)

func main() {
	store, err := db.NewStore()

	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)

	http.ListenAndServe(":3000", h)
}
