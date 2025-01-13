package main

import (
    "propertyAPI/helpers"
    _ "propertyAPI/routers"

    "github.com/beego/beego/v2/server/web"
)

func main() {
    helpers.InitDB()
    web.Run()
}