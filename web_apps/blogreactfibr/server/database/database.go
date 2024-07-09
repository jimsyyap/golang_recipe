package database

import (
    "github.com/jimsyyap/blog/model"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    _ "gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=jim password=whatsimportantnow dbname=blogreactfibr port=5432 sslmode=disable TimeZone=Australia/Melbourne"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    // https://github.com/go-gorm/postgres
    //db, err := gorm.Open(postgres.New(postgres.Config{
    //DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
    //PreferSimpleProtocol: true, // disables implicit prepared statement usage
    //}), &gorm.Config{})

    /*
    dsn := "host=localhost user=jim password=whatsimportantnow dbname=blogreactfibr port=5432 sslmode=disable TimeZone=Australia/Melbourne"
    //dsn := "root:jim@tcp(localhost:5432)/blogreactfibr?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Error),
    })
    */

    if err != nil {
        panic("failed to connect database")
    }
    log.Println("Connection Opened to Database")

    db.AutoMigrate(new(model.Blog))

    DBConn = db
}

/*
sqlDB, err := sql.Open("pgx", "blogreactfibr")
gormDB, err := gorm.Open(postgres.New(postgres.Config{
  Conn: sqlDB,
}), &gorm.Config{})
*/
