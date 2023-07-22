package handlers

import (
	"net/http"

	"github.com/jjdechavez/homelobby/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var homeView *views.View
	homeView = views.NewView("app", "views/home.html")
	homeView.Render(w, map[string]interface{}{"title": "Dashboard", "msg": "hello world"})
}
