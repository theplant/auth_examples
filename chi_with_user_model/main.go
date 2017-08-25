package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/middlewares"
)

var db, _ = gorm.Open("sqlite3", "test.db")

type User struct {
	gorm.Model
	Email string
	Name  string
}

func init() {
	db.AutoMigrate(&User{})
}

func main() {
	Auth := clean.New(&auth.Config{
		DB:        db,
		UserModel: &User{},
	})

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(
			"<html><body>Welcome %#v<br><a href='/auth/login'>Login</a><br><a href='/auth/logout'>Logout</a></body></html>", Auth.GetCurrentUser(r),
		)))
	})

	r.Mount("/auth/", Auth.NewServeMux())

	fmt.Println("Listening on: 3000")
	http.ListenAndServe(":3000", middlewares.Apply(r))
}
