package models

type PropertyList struct {
    Value     string `json:"value"`
    DestType  string `json:"dest_type"`
    HotelID   string `json:"hotel_id"`
    HotelName string `json:"hotel_name"`
    Location  string `json:"location"`
    Type      string `json:"type"`
}

type PropertyDetail struct {
    HotelID     string   `json:"hotel_id"`
    HotelName   string   `json:"hotel_name"`
    Location    string   `json:"location"`
    Rating      float64  `json:"rating"`
    ReviewCount int      `json:"review_count"`
    Price       string   `json:"price"`
    Bedrooms    int      `json:"bedrooms"`
    Bathroom    int      `json:"bathroom"`
    GuestCount  int      `json:"guest_count"`
    Description string   `json:"description"`
    ImageURLs   []string `json:"image_urls"`
    Type        string   `json:"type"`
    Amenities   []string `json:"amenities"`
}