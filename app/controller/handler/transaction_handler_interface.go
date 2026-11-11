package handler

import "net/http"

type TransactionHandlerInterface interface {
	TransactionCreate(w http.ResponseWriter, r *http.Request)
	TransactionGetAll(w http.ResponseWriter, r *http.Request)
	TransactionGetByID(w http.ResponseWriter, r *http.Request)
	TransactionUpdate(w http.ResponseWriter, r *http.Request)
	TransactionDelete(w http.ResponseWriter, r *http.Request)
}
