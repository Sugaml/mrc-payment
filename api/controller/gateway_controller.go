package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Sugaml/mrc-payment/api/models"
	"github.com/Sugaml/mrc-payment/api/repository"
	"github.com/Sugaml/mrc-payment/api/responses"
	khalti "github.com/Sugaml/mrc-payment/api/utils/Khalti"
)

var gRepo = repository.NewGetwayRepo()
var sRepo = repository.NewStudentRepo()

// CreateGateway godoc
// @Summary Create a new Gateway
// @Description Create a new Gateway with the input payload
// @Tags Gateway
// @Accept  json
// @Produce  json
// @Param body body doc.Gateway true "Create Gateway"
// @Success 201 {object} doc.Gateway
// @Router /payment/gateway [post]
func (server *Server) CreateGateway(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	data := models.Gateway{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataCreated, err := gRepo.CreateGateway(server.DB, data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

func (s *Server) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	sid, err := strconv.ParseUint(r.Header.Get("x-student-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student, err := sRepo.FindbyId(s.DB, uint(sid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	initRequest, err := khalti.NewInitiateTransactionRequest(student)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, initRequest)
}

func (s *Server) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()
	data := khalti.VerifyPaymentRequest{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// student, err := sRepo.FindbyId(s.DB, uint(data.SID))
	// if err != nil {
	// 	responses.ERROR(w, http.StatusNotFound, err)
	// 	return
	// }
	// fmt.Println(data)
	verifyPayment, err := khalti.VerifyPayment(data.Token, data.Amount)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tdata := models.Transaction{
		Title: fmt.Sprintf("Fee Payment on %s , %s", time.Now().Month().String(), fmt.Sprint(time.Now().Date())),
		//UserID:      student.ID,
		Amount:      float64(verifyPayment.State.Amount),
		GatewayID:   1,
		RefrenceID:  verifyPayment.TransactionID,
		PaymentMode: "Wallet",
		PayLoad:     fmt.Sprint(verifyPayment),
		Status:      "Completed",
	}
	transaction, err := tRepo.CreateTransaction(s.DB, tdata)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, transaction)
}
