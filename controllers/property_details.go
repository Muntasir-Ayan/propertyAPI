package controllers

import (
    "strings"
    "propertyAPI/helpers"
    "propertyAPI/models"
    "github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
    web.Controller
}

func (c *PropertyDetailsController) Get() {
    query := `
        SELECT h.hotel_id, h.hotel_name, h.location, h.rating, h.review_count, h.price, 
               h.bedrooms, h.bathroom, h.guest_count, pd.description, pd.image_url, 
               pd.type, pd.amenities
        FROM associate_hotel h
        JOIN property_detail pd ON h.hotel_id = pd.hotel_id
    `

    rows, err := helpers.DB.Query(query)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }
    defer rows.Close()

    var details []models.PropertyDetail
    for rows.Next() {
        var d models.PropertyDetail
        var imageURLArray, amenitiesArray string
        err := rows.Scan(&d.HotelID, &d.HotelName, &d.Location, &d.Rating, &d.ReviewCount, 
                         &d.Price, &d.Bedrooms, &d.Bathroom, &d.GuestCount, &d.Description, 
                         &imageURLArray, &d.Type, &amenitiesArray)
        if err != nil {
            c.Data["json"] = map[string]interface{}{
                "error": err.Error(),
            }
            c.ServeJSON()
            return
        }

        // Convert from PostgreSQL array format to Go slice
        d.ImageURLs = strings.Split(imageURLArray[1:len(imageURLArray)-1], ",")
        d.Amenities = strings.Split(amenitiesArray[1:len(amenitiesArray)-1], ",")

        details = append(details, d)
    }

    c.Data["json"] = details
    c.ServeJSON()
}