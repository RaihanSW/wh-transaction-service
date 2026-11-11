package handler

import "net/http"

type TransactionStatusHandlerInterface interface {
	TransactionStatusCreate(w http.ResponseWriter, r *http.Request)
	TransactionStatusGetAll(w http.ResponseWriter, r *http.Request)
	TransactionStatusGetByID(w http.ResponseWriter, r *http.Request)
	TransactionStatusUpdate(w http.ResponseWriter, r *http.Request)
	TransactionStatusDelete(w http.ResponseWriter, r *http.Request)
}
