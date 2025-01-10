package models

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/beego/beego/v2/server/web"
    _ "github.com/lib/pq"
)

type PropertyList struct {
    Value       string  `json:"value"`
    HotelName   string  `json:"hotel_name"`
    HotelID     string  `json:"hotel_id"`
    Rating      float64 `json:"rating"`
    ReviewCount int     `json:"review_count"`
    Type        string  `json:"type"`
}

type PropertyDetail struct {
    HotelID     string   `json:"hotel_id"`
    HotelName   string   `json:"hotel_name"`
    Description string   `json:"description"`
    ImageURLs   []string `json:"image_url"`
}

var db *sql.DB

func init() {
    dbhost, err := web.AppConfig.String("database::pg_host")
    if err != nil {
        log.Fatalf("Error getting pg_host: %s", err)
    }

    dbport, err := web.AppConfig.String("database::pg_port")
    if err != nil {
        log.Fatalf("Error getting pg_port: %s", err)
    }

    dbuser, err := web.AppConfig.String("database::pg_user")
    if err != nil {
        log.Fatalf("Error getting pg_user: %s", err)
    }

    dbpassword, err := web.AppConfig.String("database::pg_password")
    if err != nil {
        log.Fatalf("Error getting pg_password: %s", err)
    }

    dbname, err := web.AppConfig.String("database::pg_db")
    if err != nil {
        log.Fatalf("Error getting pg_db: %s", err)
    }

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbhost, dbport, dbuser, dbpassword, dbname)

    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database: %s", err)
    }

    // Verify the connection with a ping
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %s", err)
    }
}

func GetPropertyList() ([]PropertyList, error) {
    query := `
        SELECT l.value, h.hotel_name, h.hotel_id, h.rating, h.review_count, 
               pd.type
        FROM locations l
        JOIN associate_hotel h ON l.dest_id = h.dest_id
        JOIN property_detail pd ON h.hotel_id = pd.hotel_id
    `

    rows, err := db.Query(query)
    if err != nil {
        log.Printf("Error querying database: %s", err)
        return nil, err
    }
    defer rows.Close()

    var properties []PropertyList
    for rows.Next() {
        var p PropertyList
        err := rows.Scan(&p.Value, &p.HotelName, &p.HotelID, &p.Rating, 
                        &p.ReviewCount, &p.Type)
        if err != nil {
            log.Printf("Error scanning row: %s", err)
            return nil, err
        }

        properties = append(properties, p)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        log.Printf("Error iterating rows: %s", err)
        return nil, err
    }

    log.Printf("Successfully retrieved properties: %+v", properties)
    return properties, nil
}

func GetPropertyDetails() ([]PropertyDetail, error) {
    query := `
        SELECT h.hotel_id, h.hotel_name, pd.description, pd.image_url
        FROM associate_hotel h
        JOIN property_detail pd ON h.hotel_id = pd.hotel_id
    `

    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var details []PropertyDetail
    for rows.Next() {
        var d PropertyDetail
        var imageURLArray []byte
        err := rows.Scan(&d.HotelID, &d.HotelName, &d.Description, &imageURLArray)
        if err != nil {
            return nil, err
        }
        details = append(details, d)
    }

    return details, nil
}