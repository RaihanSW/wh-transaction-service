package handler

import (
	httpresponse "dc-transaction-service/app/controller/handler/httpResponse"
	"dc-transaction-service/app/model/request"
	"dc-transaction-service/app/model/response"
	"dc-transaction-service/app/service"
	"dc-transaction-service/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type transactionStatusHandler struct {
	transactionStatusService service.TransactionStatusServiceInterface
}

func NewTransactionStatusHandler(transactionStatusService service.TransactionStatusServiceInterface) TransactionStatusHandlerInterface {
	return &transactionStatusHandler{
		transactionStatusService: transactionStatusService,
	}
}

// TransactionStatusCreate handles create transaction status
// @Summary create transaction status
// @Description create transaction status
// @Tags transaction-status
// @Produce json
// @Param	TransactionStatusRequest body  request.TransactionStatusRequest true "Create Transaction Status Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /transaction/status [post]
func (h *transactionStatusHandler) TransactionStatusCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionStatusDataResponse
	var req request.TransactionStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	if status, err := util.ValidateRequest(r.Context(), req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	data, status, err := h.transactionStatusService.TransactionStatusCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// TransactionStatusGetAll handles list transaction statuses
// @Summary list transaction statuses
// @Description list all transaction statuses
// @Tags transaction-status
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /transaction/status [get]
func (h *transactionStatusHandler) TransactionStatusGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionStatusListResponse

	data, status, err := h.transactionStatusService.TransactionStatusGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// TransactionStatusGetByID handles get transaction status by id
// @Summary get transaction status by id
// @Description get transaction status by id
// @Tags transaction-status
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /transaction/status/{id} [get]
func (h *transactionStatusHandler) TransactionStatusGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionStatusDataResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	data, status, err := h.transactionStatusService.TransactionStatusGetByID(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// TransactionStatusUpdate handles update transaction status by id
// @Summary update transaction status by id
// @Description update transaction status by id
// @Tags transaction-status
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	TransactionStatusRequest body  request.TransactionStatusRequest true "Update Transaction Status Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /transaction/status/{id} [put]
func (h *transactionStatusHandler) TransactionStatusUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionStatusResponse
	var req request.TransactionStatusRequest

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	if status, err := util.ValidateRequest(r.Context(), req); err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	status, err := h.transactionStatusService.TransactionStatusUpdate(r.Context(), id, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// TransactionStatusDelete handles delete transaction status by id
// @Summary delete transaction status by id
// @Description delete transaction status by id
// @Tags transaction-status
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /transaction/status/{id} [delete]
func (h *transactionStatusHandler) TransactionStatusDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionStatusResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	status, err := h.transactionStatusService.TransactionStatusDelete(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
