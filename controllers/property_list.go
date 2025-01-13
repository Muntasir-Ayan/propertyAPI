package controllers

import (
    "log"
    "propertyAPI/helpers"
    "propertyAPI/models"

    "github.com/beego/beego/v2/server/web"
)

type PropertyListController struct {
    web.Controller
}

func (c *PropertyListController) Get() {
    query := `
        SELECT l.dest_id, l.value, l.dest_type, h.hotel_id, h.hotel_name, h.location, pd.type
        FROM locations l
        JOIN associate_hotel h ON l.dest_id = h.dest_id
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

    propertiesMap := make(map[string][]models.PropertyList)
    for rows.Next() {
        var p models.PropertyList
        err := rows.Scan(&p.DestID, &p.Value, &p.DestType, &p.HotelID, &p.HotelName, &p.Location, &p.Type)
        if err != nil {
            log.Printf("Error scanning row: %s", err)
            c.Data["json"] = map[string]interface{}{
                "error": err.Error(),
            }
            c.ServeJSON()
            return
        }
        propertiesMap[p.DestID] = append(propertiesMap[p.DestID], p)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error iterating rows: %s", err)
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    log.Printf("Successfully retrieved properties: %+v", propertiesMap)
    c.Data["json"] = propertiesMap
    c.ServeJSON()
}