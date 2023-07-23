package handlers

import (
	"net/http"

	"github.com/jjdechavez/homelobby/views"
)

func PaymentsHandler(w http.ResponseWriter, r *http.Request) {
	var paymentsView *views.View
	paymentsView = views.InitView(r, "app", "views/payments.html")
	paymentsView.Render(w, map[string]interface{}{"name": "Payments", "msg": "hello world"})
}
