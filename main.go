package main

import (
	"bmacharia/datree_room_booking_api/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	log.Println(".env file loaded successfully")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.InitDb()
}

func serveApplication() {
	router := gin.Default()
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
