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

type goodsHandler struct {
	goodsService service.GoodsServiceInterface
}

func NewGoodsHandler(goodsService service.GoodsServiceInterface) GoodsHandlerInterface {
	return &goodsHandler{
		goodsService: goodsService,
	}
}

// GoodsCreate handles create goods
// @Summary create goods
// @Description create goods
// @Tags goods
// @Produce json
// @Param	GoodsRequest body  request.GoodsRequest true "Create Goods Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /goods [post]
func (h *goodsHandler) GoodsCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsDataResponse
	var req request.GoodsRequest

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

	data, status, err := h.goodsService.GoodsCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsGetAll handles list goods
// @Summary list goods
// @Description list all goods
// @Tags goods
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /goods [get]
func (h *goodsHandler) GoodsGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsListResponse

	data, status, err := h.goodsService.GoodsGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsGetByID handles get goods by uid
// @Summary get goods by uid
// @Description get goods by uid
// @Tags goods
// @Produce json
// @Param	uid	path	string	true "Goods UID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/{uid} [get]
func (h *goodsHandler) GoodsGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsDataResponse

	uid := chi.URLParam(r, "uid")

	data, status, err := h.goodsService.GoodsGetByID(r.Context(), uid)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsUpdate handles update goods by uid
// @Summary update goods by uid
// @Description update goods by uid
// @Tags goods
// @Produce json
// @Param	uid	path	string	true "Goods UID"
// @Param	GoodsRequest body  request.GoodsRequest true "Update Goods Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/{uid} [put]
func (h *goodsHandler) GoodsUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsResponse
	var req request.GoodsRequest

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

	status, err := h.goodsService.GoodsUpdate(r.Context(), uid, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// GoodsDelete handles delete goods by uid
// @Summary delete goods by uid
// @Description delete goods by uid
// @Tags goods
// @Produce json
// @Param	uid	path	string	true "Goods UID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/{uid} [delete]
func (h *goodsHandler) GoodsDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsResponse

	uid := chi.URLParam(r, "uid")

	status, err := h.goodsService.GoodsDelete(r.Context(), uid)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
