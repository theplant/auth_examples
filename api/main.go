package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
	"github.com/qor/middlewares"
	appkitlog "github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"
)

type Order struct {
	gorm.Model
	Num   string
	State string
	Price float64
}

type User struct {
	gorm.Model
	Name string
}

func (user User) DisplayName() string {
	return user.Name
}

func main() {
	DB.DropTable(&Order{})
	DB.AutoMigrate(&User{}, &Order{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})

	Admin.SetAuth(&APIAuth{})
	Admin.AddResource(&Order{})

	mux := http.NewServeMux()
	Admin.MountTo("/api", mux)
	mux.Handle("/auth/", Auth.NewServeMux())
	color.Green("URL: %v", "http://localhost:3000/api/orders")
	server.ListenAndServe(server.Config{Addr: ":3000"}, appkitlog.Default(), middlewares.Apply(mux))
}
