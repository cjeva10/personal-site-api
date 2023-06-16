package database

import (
    "fmt"
    "log"
    "strconv"

    "git.chrisevanko.com/personal-site-api.git/config" 
    "git.chrisevanko.com/personal-site-api.git/internal/model"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    p := config.Config("DB_PORT")
    port, err := strconv.ParseUint(p, 10, 32)

    if err != nil {
        log.Println("Idiot")
    }

    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

    DB, err = gorm.Open(postgres.Open(dsn))
    if err != nil {
        panic("Failed to connect to database")
    }

    fmt.Println("Connection Opened to Database")

    DB.AutoMigrate(&model.Note{})
    DB.AutoMigrate(&model.Asset{})
    DB.AutoMigrate(&model.Price{})
    DB.AutoMigrate(&model.Exchange{})
    DB.AutoMigrate(&model.User{})
    DB.AutoMigrate(&model.ApiKey{})

    fmt.Println("Database Migrated")
}

