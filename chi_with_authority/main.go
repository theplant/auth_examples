package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/middlewares"
)

var (
	db, _     = gorm.Open("sqlite3", "test.db")
	Auth      = clean.New(&auth.Config{DB: db})
	Authority = authority.New(&authority.Config{
		Auth: Auth,
		RedirectPathAfterAccessDenied: "/auth/login",
	})
)

func main() {
	Authority.Register("logged_in_half_hour", authority.Rule{
		TimeoutSinceLastLogin: time.Minute * 30,
	})

	Authority.Register("distracted_less_than_one_minute_since_last_login", authority.Rule{
		LongestDistractionSinceLastLogin: time.Minute,
	})

	Authority.Register("logged_in_ten_minutes_and_distracted_less_than_30_seconds", authority.Rule{
		TimeoutSinceLastLogin:            time.Minute * 10,
		LongestDistractionSinceLastLogin: time.Second * 30,
	})

	r := chi.NewRouter()

	r.Get("/", defaultHandler)

	r.With(Authority.Authorize()).
		Get("/account", defaultHandler)

	r.With(Authority.Authorize("logged_in_half_hour")).
		Get("/account/edit_profile", defaultHandler)

	r.With(Authority.Authorize("distracted_less_than_one_minute_since_last_login")).
		Get("/account/edit_order", defaultHandler)

	r.With(Authority.Authorize("logged_in_ten_minutes_and_distracted_less_than_30_seconds")).
		Get("/account/edit_creditcard", defaultHandler)

	r.Mount("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 3000")
	http.ListenAndServe(":3000", middlewares.Apply(r))
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	links := []string{"<a href='/'>Home Page</a>", "<a href='/account'>Account Page</a>", "<a href='/account/edit_profile'>Edit Profile</a>", "<a href='/account/edit_order'>Edit Order</a>", "<a href='/account/edit_creditcard'>Edit Credit Card</a>"}

	claims, _ := Auth.Get(req)

	content := fmt.Sprintf("<html><body>Logged at: %v, Longest Distraction: %v<br><br>Current path: %v<br><br> Available Routers: <br>%v</body></html>", claims.LastLoginAt, claims.LongestDistractionSinceLastLogin, req.URL.Path, strings.Join(links, "<br>"))
	w.Write([]byte(content))
}
