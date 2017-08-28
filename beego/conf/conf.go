package conf

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/session"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/session/beego_session"
)

var (
	db, _     = gorm.Open("sqlite3", "test.db")
	Auth      *auth.Auth
	Authority *authority.Authority
)

func init() {
	config := `{"cookieName":"gosessionid","enableSetCookie":true,"gclifetime":3600,"ProviderConfig":"{\"cookieName\":\"gosessionid\",\"securityKey\":\"beegocookiehashkey\"}"}`
	conf := new(session.ManagerConfig)
	if err := json.Unmarshal([]byte(config), conf); err != nil {
		panic(fmt.Sprintf("json decode error: %v", err))
	}

	globalSessions, _ := session.NewManager("memory", conf)
	go globalSessions.GC()

	engine := beego_session.New(globalSessions)

	Auth = clean.New(&auth.Config{
		DB: db,
		SessionStorer: &auth.SessionStorer{
			SessionName:    "_auth_session",
			SessionManager: engine,
			SigningMethod:  jwt.SigningMethodHS256,
		},
	})

	Authority = authority.New(&authority.Config{
		Auth: Auth,
		RedirectPathAfterAccessDenied: "/auth/login",
	})
}
