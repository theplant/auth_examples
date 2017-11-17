package main

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/qor"
)

var (
	DB, _ = gorm.Open("sqlite3", "test.db")
	Auth  = clean.New(&auth.Config{
		DB:        DB,
		UserModel: User{},
	})
)

type APIAuth struct {
}

func (APIAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (APIAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (APIAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	return Auth.GetCurrentUser(c.Request).(qor.CurrentUser)
}
