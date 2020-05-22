package controllers

import (
	"net/http"
)

func (a *App) homePageGet(w http.ResponseWriter, r *http.Request) {

	// this is the home route
	http.ServeFile(w, r, "views/form.html")

}

func (a *App) homePagePost(w http.ResponseWriter, r *http.Request) {

	// this is the home route
	http.ServeFile(w, r, "views/form.html")

}
