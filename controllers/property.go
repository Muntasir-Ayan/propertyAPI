package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "propertyAPI/models"
)

type PropertyController struct {
    web.Controller
}

func (c *PropertyController) GetPropertyList() {
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

func (c *PropertyController) GetPropertyDetails() {
    details, err := models.GetPropertyDetails()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
    } else {
        c.Data["json"] = details
    }
    c.ServeJSON()
}