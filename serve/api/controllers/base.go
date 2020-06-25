package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"golang-test/serve/api/middlewares"
	"golang-test/serve/api/models"
	"golang-test/serve/api/responses"
)

// App export hear
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize connect to the database and wire up routes
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s\n", DbName)
	}

	a.DB.Debug().AutoMigrate(&models.User{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)

	a.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."+"/static/"))))

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/", home).Methods("POST")

	//login register
	a.Router.HandleFunc("/register", a.UserSignUp).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")

	//homepage
	a.Router.HandleFunc("/home", a.HomePageGet).Methods("GET")

	//action push data
	a.Router.HandleFunc("/products", a.ProductsPagePost).Methods("POST")
	a.Router.HandleFunc("/blogs", a.BlogPageGet).Methods("GET")
	a.Router.HandleFunc("/pages", a.PageGet).Methods("POST")
	a.Router.HandleFunc("/articles", a.ArticlePageGet).Methods("POST")

}

// RunServer export
func (a *App) RunServer() {
	log.Printf("\nServer starting on port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To serve")
}
