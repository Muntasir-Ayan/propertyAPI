package list_service

import (
    "log"
    "strings"
    "propertyAPI/helpers"
    "propertyAPI/models"
)

// GetPropertyList retrieves the property list from the database
func GetPropertyList() (map[string]map[string]map[string][]models.PropertyList, error) {
    query := `
        SELECT l.dest_id, l.value, l.dest_type, h.hotel_id, h.hotel_name,
        h.location, pd.type, h.rating, h.bedrooms, h.bathroom, h.guest_count,
        h.review_count, h.price, pd.description, pd.image_url, pd.amenities
        FROM locations l
        JOIN associate_hotel h ON l.dest_id = h.dest_id
        JOIN property_detail pd ON h.hotel_PropertyDetailid = pd.hotel_id
    `

    rows, err := helpers.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    propertiesMap := make(map[string]map[string]map[string][]models.PropertyList)
    for rows.Next() {
        var p models.PropertyList
        var imageURLArray, amenitiesArray string
        err := rows.Scan(&p.DestID, &p.Value, &p.DestType, &p.HotelID, 
            &p.HotelName, &p.Location, &p.Type, &p.Rating, &p.Bedrooms, 
            &p.Bathroom, &p.GuestCount, &p.ReviewCount, &p.Price, &p.Description, 
            &imageURLArray, &amenitiesArray)
        if err != nil {
            log.Printf("Error scanning row: %s", err)
            return nil, err
        }

        // Convert from PostgreSQL array format to Go slice
        p.ImageURLs = strings.Split(imageURLArray[1:len(imageURLArray)-1], ",")
        p.Amenities = strings.Split(amenitiesArray[1:len(amenitiesArray)-1], ",")

        if _, exists := propertiesMap[p.DestID]; !exists {
            propertiesMap[p.DestID] = make(map[string]map[string][]models.PropertyList)
        }
        if _, exists := propertiesMap[p.DestID][p.Location]; !exists {
            propertiesMap[p.DestID][p.Location] = make(map[string][]models.PropertyList)
        }
        propertiesMap[p.DestID][p.Location][p.HotelID] = append(propertiesMap[p.DestID][p.Location][p.HotelID], p)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error iterating rows: %s", err)
        return nil, err
    }

    log.Printf("Successfully retrieved properties: %+v", propertiesMap)
    return propertiesMap, nil
}