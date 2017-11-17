package main

import (
	scssession "github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/memstore"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/qor"
	"github.com/qor/session/manager"
	"github.com/qor/session/scs"
)

var DB *gorm.DB
var Auth *auth.Auth

func init() {
	manager.SessionManager = scs.New(scssession.NewManager(memstore.New(0)))
	DB, _ = gorm.Open("sqlite3", "test.db")
	Auth = clean.New(&auth.Config{
		DB:        DB,
		UserModel: User{},
	})
}

type APIAuth struct {
}

func (APIAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (APIAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (APIAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser, _ := Auth.GetCurrentUser(c.Request).(qor.CurrentUser)
	return currentUser
}
