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

type goodsStatusHandler struct {
	goodsStatusService service.GoodsStatusServiceInterface
}

func NewGoodsStatusHandler(goodsStatusService service.GoodsStatusServiceInterface) GoodsStatusHandlerInterface {
	return &goodsStatusHandler{
		goodsStatusService: goodsStatusService,
	}
}

// GoodsStatusCreate handles create goods status
// @Summary create goods status
// @Description create goods status
// @Tags goods-status
// @Produce json
// @Param	GoodsStatusRequest body  request.GoodsStatusRequest true "Create Goods Status Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /goods/status [post]
func (h *goodsStatusHandler) GoodsStatusCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsStatusDataResponse
	var req request.GoodsStatusRequest

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

	data, status, err := h.goodsStatusService.GoodsStatusCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsStatusGetAll handles list goods statuses
// @Summary list goods statuses
// @Description list all goods statuses
// @Tags goods-status
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /goods/status [get]
func (h *goodsStatusHandler) GoodsStatusGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsStatusListResponse

	data, status, err := h.goodsStatusService.GoodsStatusGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsStatusGetByID handles get goods status by id
// @Summary get goods status by id
// @Description get goods status by id
// @Tags goods-status
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/status/{id} [get]
func (h *goodsStatusHandler) GoodsStatusGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsStatusDataResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	data, status, err := h.goodsStatusService.GoodsStatusGetByID(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsStatusUpdate handles update goods status by id
// @Summary update goods status by id
// @Description update goods status by id
// @Tags goods-status
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	GoodsStatusRequest body  request.GoodsStatusRequest true "Update Goods Status Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/status/{id} [put]
func (h *goodsStatusHandler) GoodsStatusUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsStatusResponse
	var req request.GoodsStatusRequest

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

	status, err := h.goodsStatusService.GoodsStatusUpdate(r.Context(), id, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// GoodsStatusDelete handles delete goods status by id
// @Summary delete goods status by id
// @Description delete goods status by id
// @Tags goods-status
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/status/{id} [delete]
func (h *goodsStatusHandler) GoodsStatusDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsStatusResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	status, err := h.goodsStatusService.GoodsStatusDelete(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
