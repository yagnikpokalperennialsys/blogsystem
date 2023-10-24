package main

import (
	"backend/internal/controller"
	"backend/internal/routes"
	appconst "backend/pkg/appconstant"
	"backend/pkg/db"
	"backend/pkg/repository/dbrepo"
	services "backend/services/articles"
	"time"

	_ "github.com/lib/pq"

	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set application config
	var app routes.Application

	// Read from the command line
	flag.StringVar(&app.DSN, "dsn", "host=postgres port=5432 user=postgres password=postgres dbname=articles sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()
	fmt.Println(appconst.DatabseWait)
	time.Sleep(5 * time.Second)

	// Connect to the database
	conn, err := db.ConnectToDB(app.DSN)
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	// Create the table if it does not exist in the container
	app.DB.CreateTable()

	// Initialize the ArticleService with the DatabaseRepo
	articleService := services.NewArticleService(app.DB)

	// Create the MyApplication instance and pass the dependencies
	myApp := controller.Controller{
		DSN:            app.DSN,
		DB:             app.DB,
		Utility:        app.Utility, // You can replace this with your actual utility implementation
		ArticleService: articleService,
	}

	// Set the handlers for your application
	app.Handler = myApp

	log.Println(appconst.Startapp, appconst.Port)

	// Start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", appconst.Port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
