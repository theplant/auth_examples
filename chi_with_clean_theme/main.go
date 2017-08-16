package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
)

var db, _ = gorm.Open("sqlite3", "test.db")

func main() {
	Auth := clean.New(&auth.Config{DB: db})
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 3000")
	http.ListenAndServe(":3000", r)
}
