package menu

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services Service
}

func NewHandler(service Service) Handler {
	return Handler{
		services: service,
	}
}

func (h Handler) Create(rw http.ResponseWriter, r *http.Request) {
	var req CreateMenuRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	menu := New(req.Name, req.Category, req.Desc, req.ImageUrl, req.Price)

	err = h.services.Save(menu)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}
	WriteJsonResponse(rw, resp)
}

func (h Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	menus, err := h.services.GetAll()
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: menus,
	}
	WriteJsonResponse(rw, resp)
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

	menu, err := h.services.GetById(idInt)
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
		Payload: menu,
	}
	WriteJsonResponse(rw, resp)
}

func (h Handler) Update(rw http.ResponseWriter, r *http.Request) {
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

	var req UpdateMenuRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// Generate error response
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}

		WriteJsonResponse(rw, resp)
		return
	}

	menu := New(req.Name, req.Category, req.Desc, req.ImageUrl, req.Price).WithId(idInt)

	err = h.services.UpdateById(menu)
	if err != nil {
		// Generate error response
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
	}
	WriteJsonResponse(rw, resp)
}

func (h Handler) Delete(rw http.ResponseWriter, r *http.Request) {
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

	err = h.services.DeleteById(idInt)
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
				Message: "ERR SERVER",
				Error:   err.Error(),
			}
		}

		WriteJsonResponse(rw, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
	}

	WriteJsonResponse(rw, resp)
}
