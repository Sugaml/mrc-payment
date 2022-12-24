package controller

import (
	"net/http"

	"github.com/Sugaml/mrc-payment/api/middleware"
	"github.com/Sugaml/mrc-payment/api/responses"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

func (server *Server) setAdmin(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.setJSON(path, middleware.SetAdminMiddlewareAuthentication(next), method)
}

func (server *Server) initializeRoutes() {
	server.setJSON("/", server.WelcomePage, "GET")

	server.setJSON("/payment/invoice", server.CreateIvoice, "POST")
	server.setJSON("/payment/initiaterequest", server.InitiatePayment, "POST")
	server.setJSON("/payment/verify", server.VerifyPayment, "POST")

	server.setJSON("/payment/invoice", server.CreateGateway, "POST")
	server.setAdmin("/payment/transactions", server.GetTransaction, "GET")
	server.setJSON("payment/invoice", server.CreateIvoice, "POST")

	server.setJSON("payment/invoice", server.CreateGateway, "POST")
	server.setJSON("/payment/transactions", server.GetTransaction, "Get")
}

func (server *Server) WelcomePage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to mrc-api project")
}
