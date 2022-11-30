package controller

import (
	"net/http"
	"strconv"

	"github.com/Sugaml/mrc-payment/api/repository"
	"github.com/Sugaml/mrc-payment/api/responses"
)

var tRepo = repository.NewTransactionRepo()

// GetTransactions godoc
// @Summary Transactions of 01cloud for admin
// @Description Transactions of 01cloud for admin
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param x-user-role header string true "x-user-role"
// @Param status path string true "Status"
// @Param user_id path string true "User ID"
// @Param startdate path string true "Start Date"
// @Param enddate path string true "End Date"
// @Param size path string true "Size"
// @Param page path string true "Page"
// @Param search path string true "Search"
// @Param sort-column path string true "Sort-Column"
// @Param sort-Direction path string true "Sort-Direction"
// @Success 200 {object} doc.Transaction
// @Router  /payment/transaction [get]
func (server *Server) GetTransaction(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	userID, _ := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 32)
	size, _ := strconv.ParseUint(r.URL.Query().Get("size"), 10, 32)
	page, _ := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	search := r.URL.Query().Get("search")
	sortColumn := r.URL.Query().Get("sort-column")
	sortDirection := r.URL.Query().Get("sort-direction")
	startdate := r.URL.Query().Get("startdate")
	enddate := r.URL.Query().Get("endate")
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	if sortColumn == "" {
		sortColumn = "id"
	}
	if sortDirection == "" {
		sortDirection = "desc"
	}
	datas, count, err := tRepo.FilterTransactionByStatus(server.DB, userID, page, size, status, startdate, enddate, search, sortColumn, sortDirection)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, map[string]interface{}{
		"data":  datas,
		"total": count,
	})
}
