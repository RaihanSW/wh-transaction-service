package handler

import "net/http"

type GoodsHandlerInterface interface {
	GoodsCreate(w http.ResponseWriter, r *http.Request)
	GoodsGetAll(w http.ResponseWriter, r *http.Request)
	GoodsGetByID(w http.ResponseWriter, r *http.Request)
	GoodsUpdate(w http.ResponseWriter, r *http.Request)
	GoodsDelete(w http.ResponseWriter, r *http.Request)
}
