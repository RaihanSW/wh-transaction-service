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

type goodsCategoryHandler struct {
	goodsCategoryService service.GoodsCategoryServiceInterface
}

func NewGoodsCategoryHandler(goodsCategoryService service.GoodsCategoryServiceInterface) GoodsCategoryHandlerInterface {
	return &goodsCategoryHandler{
		goodsCategoryService: goodsCategoryService,
	}
}

// GoodsCategoryCreate handles create goods category
// @Summary create goods category
// @Description create goods category
// @Tags goods-category
// @Produce json
// @Param	GoodsCategoryRequest body  request.GoodsCategoryRequest true "Create Goods Category Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 201 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Router /goods/category [post]
func (h *goodsCategoryHandler) GoodsCategoryCreate(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsCategoryDataResponse
	var req request.GoodsCategoryRequest

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

	data, status, err := h.goodsCategoryService.GoodsCategoryCreate(r.Context(), req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsCategoryGetAll handles list goods categoryes
// @Summary list goods categoryes
// @Description list all goods categoryes
// @Tags goods-category
// @Produce json
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Router /goods/category [get]
func (h *goodsCategoryHandler) GoodsCategoryGetAll(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsCategoryListResponse

	data, status, err := h.goodsCategoryService.GoodsCategoryGetAll(r.Context())
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsCategoryGetByID handles get goods category by id
// @Summary get goods category by id
// @Description get goods category by id
// @Tags goods-category
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/category/{id} [get]
func (h *goodsCategoryHandler) GoodsCategoryGetByID(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsCategoryDataResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	data, status, err := h.goodsCategoryService.GoodsCategoryGetByID(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	resp.Data = data
	httpresponse.Response(w, r, resp, status)
}

// GoodsCategoryUpdate handles update goods category by id
// @Summary update goods category by id
// @Description update goods category by id
// @Tags goods-category
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	GoodsCategoryRequest body  request.GoodsCategoryRequest true "Update Goods Category Request Body"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 400 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/category/{id} [put]
func (h *goodsCategoryHandler) GoodsCategoryUpdate(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsCategoryResponse
	var req request.GoodsCategoryRequest

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

	status, err := h.goodsCategoryService.GoodsCategoryUpdate(r.Context(), id, req)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}

// GoodsCategoryDelete handles delete goods category by id
// @Summary delete goods category by id
// @Description delete goods category by id
// @Tags goods-category
// @Produce json
// @Param	id	path	int	true "Transaction Status ID"
// @Param	Authorization	header	string	false "Bearer token"
// @Success 200 {object}  httpresponse.SetResponse
// @Failure 404 {object}  httpresponse.SetResponse
// @Router /goods/category/{id} [delete]
func (h *goodsCategoryHandler) GoodsCategoryDelete(w http.ResponseWriter, r *http.Request) {

	var resp response.GoodsCategoryResponse

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, http.StatusBadRequest, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, http.StatusBadRequest)
		return
	}

	status, err := h.goodsCategoryService.GoodsCategoryDelete(r.Context(), id)
	if err != nil {
		httpresponse.ErrorHandler(r.Context(), err, status, &resp.GeneralResponse)
		httpresponse.Response(w, r, resp, status)
		return
	}

	resp.IsSuccess = true
	httpresponse.Response(w, r, resp, status)
}
