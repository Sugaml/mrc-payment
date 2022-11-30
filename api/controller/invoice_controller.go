package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sugaml/mrc-payment/api/models"
	"github.com/Sugaml/mrc-payment/api/repository"
	"github.com/Sugaml/mrc-payment/api/responses"
)

var iRepo = repository.NewInvoiceRepo()

// CreateGateway godoc
// @Summary Create a new Gateway
// @Description Create a new Gateway with the input payload
// @Tags Gateway
// @Accept  json
// @Produce  json
// @Param body body doc.Gateway true "Create Gateway"
// @Success 201 {object} doc.Gateway
// @Router /payment/gateway [post]
func (server *Server) CreateIvoice(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Invoice{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataCreated, err := iRepo.CreateInvoice(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}
