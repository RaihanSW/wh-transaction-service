package handler

import (
	httpresponse "dc-transaction-service/app/controller/handler/httpResponse"
	"dc-transaction-service/app/model/request"
	"dc-transaction-service/app/model/response"
	"dc-transaction-service/app/service"
	"dc-transaction-service/util"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type transactionHandler struct {
	transactionService service.TransactionServiceInterface
}

func NewTransactionHandler(transactionService service.TransactionServiceInterface) TransactionHandlerInterface {
	return &transactionHandler{
		transactionService: transactionService,
	}
}

// TransactionCreate handles create transaction
// @Summary create transaction
// @Description create transaction
// @Tags transaction
// @Produce json
// @Param	TransactionRequest body  request.TransactionRequest true "Create Transaction Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /transaction [post]
func (h *transactionHandler) TransactionCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionDataResponse
	var req request.TransactionRequest

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

	data, status, err := h.transactionService.TransactionCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// TransactionGetAll handles list transactions
// @Summary list transactions
// @Description list all transactions
// @Tags transaction
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /transaction [get]
func (h *transactionHandler) TransactionGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionListResponse

	data, status, err := h.transactionService.TransactionGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// TransactionGetByID handles get transaction by uid
// @Summary get transaction by uid
// @Description get transaction by uid
// @Tags transaction
// @Produce json
// @Param	uid	path	string	true "Transaction UID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /transaction/{uid} [get]
func (h *transactionHandler) TransactionGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionDataResponse

	uid := chi.URLParam(r, "uid")

	data, status, err := h.transactionService.TransactionGetByID(r.Context(), uid)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// TransactionUpdate handles update transaction by uid
// @Summary update transaction by uid
// @Description update transaction by uid
// @Tags transaction
// @Produce json
// @Param	uid	path	string	true "Transaction UID"
// @Param	TransactionRequest body  request.TransactionRequest true "Update Transaction Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /transaction/{uid} [put]
func (h *transactionHandler) TransactionUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionResponse
	var req request.TransactionRequest

	uid := chi.URLParam(r, "uid")

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

	status, err := h.transactionService.TransactionUpdate(r.Context(), uid, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// TransactionDelete handles delete transaction by uid
// @Summary delete transaction by uid
// @Description delete transaction by uid
// @Tags transaction
// @Produce json
// @Param	uid	path	string	true "Transaction UID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /transaction/{uid} [delete]
func (h *transactionHandler) TransactionDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.TransactionResponse

	uid := chi.URLParam(r, "uid")

	status, err := h.transactionService.TransactionDelete(r.Context(), uid)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
