package handler

import "net/http"

type GoodsStatusHandlerInterface interface {
	GoodsStatusCreate(w http.ResponseWriter, r *http.Request)
	GoodsStatusGetAll(w http.ResponseWriter, r *http.Request)
	GoodsStatusGetByID(w http.ResponseWriter, r *http.Request)
	GoodsStatusUpdate(w http.ResponseWriter, r *http.Request)
	GoodsStatusDelete(w http.ResponseWriter, r *http.Request)
}
