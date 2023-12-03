package auth

import (
	"encoding/json"
	"net/http"
	routerChi "nobee-pos/infra/router/chi"
	"nobee-pos/utility"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := routerChi.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		routerChi.WriteJsonResponse(w, resp)
		return
	}

	auth := New(req.Email, req.Password)

	err = h.svc.Create(auth)
	if err != nil {
		resp := routerChi.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}
		routerChi.WriteJsonResponse(w, resp)
		return
	}
	resp := routerChi.APIResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}
	routerChi.WriteJsonResponse(w, resp)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := routerChi.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		routerChi.WriteJsonResponse(w, resp)
		return
	}

	auth := New(req.Email, req.Password)

	newAuth, err := h.svc.Login(auth)
	if err != nil {
		resp := routerChi.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}
		routerChi.WriteJsonResponse(w, resp)
		return
	}

	token := utility.NewJWT(newAuth.Id)

	tokString, err := token.GenerateToken()
	if err != nil {
		resp := routerChi.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}
		routerChi.WriteJsonResponse(w, resp)
		return
	}

	resp := routerChi.APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",

		Payload: map[string]interface{}{
			"token": tokString,
		},
	}
	routerChi.WriteJsonResponse(w, resp)
}

func (h Handler) Profile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("AUTH_ID")

	if id == nil {
		resp := routerChi.APIResponse{
			Status:  http.StatusForbidden,
			Message: "FORBIDDEN ACCESS",
		}
		routerChi.WriteJsonResponse(w, resp)
		return
	}

	resp := routerChi.APIResponse{
		Status:  http.StatusOK,
		Message: "GET PROFILE",
		Payload: id,
	}

	routerChi.WriteJsonResponse(w, resp)
}
