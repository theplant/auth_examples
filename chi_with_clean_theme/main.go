package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
)

func main() {
	Auth := clean.New(&auth.Config{})
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 3000")
	http.ListenAndServe(":3000", r)
}
