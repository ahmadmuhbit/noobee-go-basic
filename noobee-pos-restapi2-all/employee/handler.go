package employee

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

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

func (h Handler) ListEmployees(w http.ResponseWriter, r *http.Request) {
	emp, err := h.svc.FindAll()
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}

		WriteJsonResponse(w, resp)
		return
	}

	if len(emp) == 0 {
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
		Message: "Success",
		Payload: emp,
	}

	WriteJsonResponse(w, resp)
}

func (h Handler) Add(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}

		WriteJsonResponse(w, resp)
		return
	}

	emp := New(req.NIP, req.Name, req.Address)
	err = h.svc.Create(emp)
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
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}

	WriteJsonResponse(w, resp)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   "Id is required",
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

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}

		WriteJsonResponse(w, resp)
		return
	}

	emp := New(req.NIP, req.Name, req.Address)
	err = h.svc.UpdateById(emp, idInt)
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
	}

	WriteJsonResponse(w, resp)
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   "Id is required",
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

	err = h.svc.DeleteById(idInt)
	if err != nil {
		resp := APIResponse{}

		if err == sql.ErrNoRows {
			resp = APIResponse{
				Status:  http.StatusNotFound,
				Message: "ERROR NOT FOUND",
				Error:   err.Error(),
			}
		} else {
			resp = APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "ERROR SERVER",
				Error:   err.Error(),
			}
		}

		WriteJsonResponse(w, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
	}

	WriteJsonResponse(w, resp)
}

func (h Handler) GetById(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   "id is required",
		}
		WriteJsonResponse(rw, resp)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	emp, err := h.svc.GetById(idInt)
	if err != nil {
		resp := APIResponse{}

		if err == sql.ErrNoRows {
			resp = APIResponse{
				Status:  http.StatusNotFound,
				Message: "ERROR NOT FOUND",
				Error:   err.Error(),
			}
		} else {
			resp = APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "ERROR SERVER",
				Error:   err.Error(),
			}
		}
		WriteJsonResponse(rw, resp)
		return
	}
	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: emp,
	}
	WriteJsonResponse(rw, resp)
}
