package main

import (
	"github.com/joho/godotenv"
	"log"
	"pointSystem/db"
	"pointSystem/docs"
	"pointSystem/route"
	"pointSystem/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	docs.SwaggerInfo.Title = "Swagger Rest API CRUD User"
	docs.SwaggerInfo.Description = "This is a CRUD User and User Login API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = util.GetEnv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	database := db.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()

	r := route.SetupRouter(database)
	r.Run()
}
