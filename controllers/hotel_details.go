// controllers/hotel_details.go
package controllers

import (
    "propertyAPI/hotel_service"
    "github.com/beego/beego/v2/server/web"
)

type HotelDetailsController struct {
    web.Controller
}

func (c *HotelDetailsController) Get() {
    hotelID := c.GetString("hotel_id")
    if hotelID == "" {
        c.Data["json"] = map[string]interface{}{
            "error": "hotel_id parameter is required",
        }
        c.ServeJSON()
        return
    }

    details, err := hotel_service.GetHotelDetailsByID(hotelID)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = details
    c.ServeJSON()
}