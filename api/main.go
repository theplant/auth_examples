package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/qor-example-cases/config"
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

func main() {
	var Admin = config.Admin

	Admin.SetAuth(&APIAuth{})
	Admin.AddResource(&Order{})

	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	color.Green("URL: %v", "http://localhost:3000/admin/orders")
	server.ListenAndServe(server.Config{Addr: ":3000"}, appkitlog.Default(), mux)
}
