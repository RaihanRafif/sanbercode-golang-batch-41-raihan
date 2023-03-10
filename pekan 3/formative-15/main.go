package main

import (
	"database/sql"
	"fmt"
	"formative-15/controllers"
	"formative-15/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	//ENV
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	// psqlInfo := fmt.Sprintf(`host=#{os.Getenv("DB_HOST")} port=#{os.Getenv("DB_PORT")} user=#{os.Getenv("DB_USER")} password=#{os.Getenv("DB_PASSWORD")} dbname=#{os.Getenv("DB_NAME")} sslmode=disable`)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	router := gin.Default()
	router.GET("/person", controllers.GetAllPerson)
	router.POST("/person", controllers.InsertPerson)
	router.PUT("/person", controllers.UpdatePerson)
	router.DELETE("/person", controllers.DeletePerson)

	router.Run("localhost:8080")
}
