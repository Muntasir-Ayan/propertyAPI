//helpers/db.go

package helpers

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/beego/beego/v2/server/web"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
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

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database: %s", err)
    }

    // Verify the connection with a ping
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %s", err)
    }
}