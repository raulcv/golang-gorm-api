package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/raulcv/golang-gorm-api/database"
	"github.com/raulcv/golang-gorm-api/models"
	"github.com/raulcv/golang-gorm-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("Environment variables loaded successfully")
	}
	port := os.Getenv("HTTP_PORT")
	// fmt.Println("port: ", port)
	// database connection
	database.DBConnection()

	// db.DB.Migrator().DropTable(models.User{})
	database.DB.AutoMigrate(models.Task{})
	database.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// Index route
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":"+port, r)
}
