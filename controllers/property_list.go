package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "propertyAPI/models"
)

type PropertyListController struct {
    web.Controller
}

func (c *PropertyListController) Get() {
    properties, err := models.GetPropertyList()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
    } else {
        c.Data["json"] = properties
    }
    c.ServeJSON()
}