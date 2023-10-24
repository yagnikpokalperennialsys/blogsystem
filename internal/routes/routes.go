package routes

import (
	"backend/internal/controller"
	"backend/pkg/repository/dbrepo"
	services "backend/services/articles"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

type Application struct {
	DSN            string
	DB             dbrepo.DatabaseRepo
	Utility        controller.UtilityInterface
	ArticleService *services.ArticleService
	Handler        controller.Controller
}

func (app *Application) Routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()
	// Create a new CORS middleware instance with your desired options.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace with allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           600,
	})
	// Use the CORS middleware
	mux.Use(c.Handler)
	mux.Use(middleware.Recoverer)
	mux.Get("/", app.Handler.HealthCheck)
	mux.Get("/articles", app.Handler.AllArticle)
	mux.Get("/articles/{id}", app.Handler.GetArticle)
	mux.Post("/articles", app.Handler.InsertArticle)

	return mux
}
