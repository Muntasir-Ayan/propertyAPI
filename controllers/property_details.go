package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "propertyAPI/models"
)

type PropertyDetailsController struct {
    web.Controller
}

func (c *PropertyDetailsController) Get() {
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