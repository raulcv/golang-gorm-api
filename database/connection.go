package database

import (
	"log"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db_host = os.Getenv("DB_HOST")
var db_user_name = os.Getenv("DB_USER")
var db_user_pwd = os.Getenv("DB_PWD")
var db_name = os.Getenv("DB_NAME")
var db_port = os.Getenv("DB_PORT")

var dbPropers = []string{"host=" + db_host, "user=" + db_user_name, "password=" + db_user_pwd,
	"dbname=" + db_name,
	"port=" + db_port}

// Data Base String = dsn
var dsn = strings.Join(dbPropers, " ")

var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("gormapi db connected")
	}
}
