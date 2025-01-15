package controllers

import (
    "propertyAPI/list_service"

    beego "github.com/beego/beego/v2/server/web"
)

type PropertyListController struct {
    beego.Controller
}

func (c *PropertyListController) Get() {
    propertiesMap, err := list_service.GetPropertyList()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = propertiesMap
    c.ServeJSON()
}