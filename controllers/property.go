package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "strings"
    "propertyAPI/models"
)

type PropertyController struct {
    web.Controller
}

func (c *PropertyController) Get() {
    // Get the requested URL path
    requestPath := c.Ctx.Request.URL.Path

    if strings.HasSuffix(requestPath, "/list") {
        c.getPropertyList()
    } else if strings.HasSuffix(requestPath, "/details") {
        c.getPropertyDetails()
    } else {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid endpoint. Use '/list' or '/details'.",
        }
        c.ServeJSON()
    }
}

func (c *PropertyController) getPropertyList() {
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

func (c *PropertyController) getPropertyDetails() {
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
