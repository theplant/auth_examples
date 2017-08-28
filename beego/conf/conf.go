package conf

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth_themes/clean"
)

var (
	db, _     = gorm.Open("sqlite3", "test.db")
	Auth      = clean.New(&auth.Config{DB: db})
	Authority = authority.New(&authority.Config{
		Auth: Auth,
		RedirectPathAfterAccessDenied: "/auth/login",
	})
)
