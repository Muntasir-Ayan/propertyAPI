// destination_service/destination_service.go
package hotel_service

import (
    "strings"
    "propertyAPI/helpers"
    "propertyAPI/models"
)

func GetHotelDetailsByID(hotelID string) (*models.PropertyDetail, error) {
    query := `
        SELECT h.hotel_id, h.hotel_name, h.location, h.rating, h.review_count, h.price, 
               h.bedrooms, h.bathroom, h.guest_count, pd.description, pd.image_url, 
               pd.type, pd.amenities
        FROM associate_hotel h
        JOIN property_detail pd ON h.hotel_id = pd.hotel_id
        WHERE h.hotel_id = $1
    `

    row := helpers.DB.QueryRow(query, hotelID)

    var d models.PropertyDetail
    var imageURLArray, amenitiesArray string
    err := row.Scan(&d.HotelID, &d.HotelName, &d.Location, &d.Rating, &d.ReviewCount, 
                    &d.Price, &d.Bedrooms, &d.Bathroom, &d.GuestCount, &d.Description, 
                    &imageURLArray, &d.Type, &amenitiesArray)
    if err != nil {
        return nil, err
    }

    // Convert from PostgreSQL array format to Go slice
    d.ImageURLs = strings.Split(imageURLArray[1:len(imageURLArray)-1], ",")
    d.Amenities = strings.Split(amenitiesArray[1:len(amenitiesArray)-1], ",")

    return &d, nil
}