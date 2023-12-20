package sql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	schema "notes-api-golang/framework/sql/schemas"
)

var SqlDB *gorm.DB

func ConnectMysql() {
	godotenv.Load()

	host := os.Getenv("SQLDB_HOST")
	port := os.Getenv("SQLDB_PORT")
	user := os.Getenv("SQLDB_USER")
	password := os.Getenv("SQLDB_PASSWORD")
	dbname := os.Getenv("SQLDB_NAME")

	// check if all env is set
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		panic("Environment variable not set")
	}

	// connect to db
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	} else {
		fmt.Println("Connected to database")
		SqlDB = db
	}

	// auto migrate
	db.AutoMigrate(&schema.User{})
}
