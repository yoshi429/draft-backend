package handler

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/application/usecase"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/request"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error
	RefreshToken(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	*request.Context
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase, c *request.Context) *userHandler {
	return &userHandler{
		Context:     c,
		userUseCase: userUseCase,
	}
}

type requestUser struct {
}

type responseUser struct {
}

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (uh *userHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, _ := auth.GenerateBcryptPassword(password)

	err := uh.userUseCase.Create(name, hash, email)
	if err != nil {
		return uh.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return uh.JSON(w, http.StatusOK, nil)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	token, err := uh.userUseCase.Login(email, password)
	if err != nil {
		return uh.Error(w, http.StatusInternalServerError, err)
	}

	res := loginResponse{
		AccessToken:  token.AccessToken.JwtToken(),
		RefreshToken: token.RefreshToken.JwtToken(),
	}
	return uh.JSON(w, http.StatusOK, res)
}

type refreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type refreshTokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *userHandler) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("hogehoge")
	var req refreshTokenReq
	err := h.Bind(r, &req)
	if err != nil {
		return h.Error(w, http.StatusInternalServerError, err)
	}

	authToken, err := h.userUseCase.RefreshToken(req.RefreshToken)
	if err != nil {
		return h.Error(w, http.StatusInternalServerError, err)
	}

	res := refreshTokenRes{
		AccessToken:  authToken.AccessToken.JwtToken(),
		RefreshToken: authToken.RefreshToken.JwtToken(),
	}

	return h.JSON(w, http.StatusOK, res)
}
