package handler

import "net/http"

type GoodsCategoryHandlerInterface interface {
	GoodsCategoryCreate(w http.ResponseWriter, r *http.Request)
	GoodsCategoryGetAll(w http.ResponseWriter, r *http.Request)
	GoodsCategoryGetByID(w http.ResponseWriter, r *http.Request)
	GoodsCategoryUpdate(w http.ResponseWriter, r *http.Request)
	GoodsCategoryDelete(w http.ResponseWriter, r *http.Request)
}
