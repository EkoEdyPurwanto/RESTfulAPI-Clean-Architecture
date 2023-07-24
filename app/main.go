package main

import (
	"EchoEdyP/RESTfulAPI-Clean-Architecture/category/delivery/http"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/category/repository/postgres"
	"EchoEdyP/RESTfulAPI-Clean-Architecture/category/usecase"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.DBname`)
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName) // Update the connection string for PostgreSQL

	dbConn, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	validate := validator.New()
	categoryRepository := postgres.NewCategoryRepository()
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, dbConn, validate)
	http.NewCategoryDelivery(e, categoryUseCase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
