package main

import (
    "propertyAPI/helpers"
    _ "propertyAPI/routers"

    beego "github.com/beego/beego/v2/server/web"
)

func main() {
    helpers.InitDB()
    beego.Run()
}