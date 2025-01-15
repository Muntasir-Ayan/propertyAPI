// destination_service/destination_service.go
package destination_service

import (
    "strings"
    "propertyAPI/helpers"
    "propertyAPI/models"
)

func GetPropertyDetailsByDestination(destID string) (map[string]interface{}, error) {
    query := `
        SELECT h.hotel_id, h.hotel_name, h.location, h.rating, h.review_count, h.price, 
               h.bedrooms, h.bathroom, h.guest_count, pd.description, pd.image_url, 
               pd.type, pd.amenities, l.value, l.dest_type
        FROM associate_hotel h
        JOIN property_detail pd ON h.hotel_id = pd.hotel_id
        JOIN locations l ON h.dest_id = l.dest_id
        WHERE h.dest_id = $1
    `

    rows, err := helpers.DB.Query(query, destID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make(map[string]interface{})

    for rows.Next() {
        var d models.PropertyDetail
        var imageURLArray, amenitiesArray string
        var location, value, destType string
        err := rows.Scan(&d.HotelID, &d.HotelName, &d.Location, &d.Rating, &d.ReviewCount, 
                         &d.Price, &d.Bedrooms, &d.Bathroom, &d.GuestCount, &d.Description, 
                         &imageURLArray, &d.Type, &amenitiesArray, &value, &destType)
        if err != nil {
            return nil, err
        }

        // Convert from PostgreSQL array format to Go slice
        d.ImageURLs = strings.Split(imageURLArray[1:len(imageURLArray)-1], ",")
        d.Amenities = strings.Split(amenitiesArray[1:len(amenitiesArray)-1], ",")

        if _, ok := result[destID]; !ok {
            result[destID] = make(map[string]interface{})
        }

        locationMap := result[destID].(map[string]interface{})
        if _, ok := locationMap[location]; !ok {
            locationMap[location] = make(map[string]interface{})
            locationMap[location] = map[string]interface{}{
                "value":     value,
                "dest_type": destType,
            }
        }

        hotelMap := locationMap[location].(map[string]interface{})
        if _, ok := hotelMap[d.HotelID]; !ok {
            hotelMap[d.HotelID] = []models.PropertyDetail{}
        }

        hotelDetails := hotelMap[d.HotelID].([]models.PropertyDetail)
        hotelDetails = append(hotelDetails, d)
        hotelMap[d.HotelID] = hotelDetails
    }

    return result, nil
}