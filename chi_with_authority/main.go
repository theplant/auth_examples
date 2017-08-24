package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/middlewares"
)

var db, _ = gorm.Open("sqlite3", "test.db")

func main() {
	var (
		Auth      = clean.New(&auth.Config{DB: db})
		Authority = authority.New(&authority.Config{
			Auth: Auth,
			RedirectPathAfterAccessDenied: "/auth/login",
		})
	)

	Authority.Register("last_actived_in_half_hour", authority.Rule{TimeoutSinceLastActive: time.Minute * 30})

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.With(Authority.Authorize()).Get("/account", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("account page"))
	})

	r.With(Authority.Authorize("last_actived_in_half_hour")).Get("/account/add_creditcard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("add credit card"))
	})

	r.Mount("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 3000")
	http.ListenAndServe(":3000", middlewares.Apply(r))
}
