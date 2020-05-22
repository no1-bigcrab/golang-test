package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"golang-first/serve/api/middlewares"
	"golang-first/serve/api/models"
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
	a.Router.HandleFunc("/home", a.homePageGet).Methods("GET")
	a.Router.HandleFunc("/home", a.homePagePost).Methods("POST")

}

// RunServer export
func (a *App) RunServer() {
	log.Printf("\nServer starting on port 9000 \n")
	log.Fatal(http.ListenAndServe(":9000", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	// this is the home route

	methods := r.Method

	fmt.Println("method:", r.Method) //get request method

	switch methods {
	case "GET":
		w.Header().Set("Content-Type", "text/html")

		tmpl := template.Must(template.ParseFiles("./api/views/form.html"))
		tmpl.Execute(w, r)

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)

	default:
		fmt.Println("don't have methods")
	}
}
