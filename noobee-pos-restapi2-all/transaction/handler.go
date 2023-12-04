package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	trxs, err := h.svc.ListAll()
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}

		WriteJsonResponse(w, resp)
		return
	}

	if len(trxs) == 0 {
		resp := APIResponse{
			Status:  http.StatusNotFound,
			Message: "DATA NOT FOUND",
			Error:   "data not found",
		}

		WriteJsonResponse(w, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: trxs,
	}

	WriteJsonResponse(w, resp)
}

func (h Handler) DetailById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   "id is required",
		}

		WriteJsonResponse(w, resp)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}

	trx, err := h.svc.Detail(idInt)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: trx,
	}
	WriteJsonResponse(w, resp)
}

func (h Handler) Pay(w http.ResponseWriter, r *http.Request) {
	var req = PayRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}

	var trx = NewFromPayRequest(req)
	err := h.svc.Pay(trx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			resp := APIResponse{
				Status:  http.StatusNotFound,
				Message: "DATA NOT FOUND",
				Error:   err.Error(),
			}
			WriteJsonResponse(w, resp)
			return
		}
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}
	resp := APIResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}
	WriteJsonResponse(w, resp)
}
