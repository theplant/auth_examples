package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/auth"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/middlewares"
)

var db, _ = gorm.Open("sqlite3", "test.db")

func main() {
	Auth := clean.New(&auth.Config{DB: db})

	Auth.RegisterProvider(github.New(&github.Config{
		ClientID:     "client id",
		ClientSecret: "client secret",
	}))

	Auth.RegisterProvider(google.New(&google.Config{
		ClientID:     "client id",
		ClientSecret: "client secret",
	}))

	Auth.RegisterProvider(facebook.New(&facebook.Config{
		ClientID:     "client id",
		ClientSecret: "client secret",
	}))

	Auth.RegisterProvider(twitter.New(&twitter.Config{
		ClientID:     "client id",
		ClientSecret: "client secret",
	}))

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 3000")
	http.ListenAndServe(":3000", middlewares.Apply(r))
}
