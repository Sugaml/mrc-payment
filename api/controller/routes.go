package controller

import (
	"net/http"

	"github.com/Sugaml/mrc-payment/api/middleware"
	"github.com/Sugaml/mrc-payment/api/responses"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

// func (server *Server) setAdmin(path string, next func(http.ResponseWriter, *http.Request), method string) {
// 	server.setJSON(path, middleware.SetAdminMiddlewareAuthentication(next), method)
// }

func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.CORS)
	server.setJSON("/", server.WelcomePage, "GET")

	server.setJSON("/invoice", server.CreateIvoice, "POST")

	server.setJSON("/invoice", server.CreateGateway, "POST")
}

func (server *Server) WelcomePage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome to mrc-api project")
}
