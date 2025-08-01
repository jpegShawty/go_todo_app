package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/jpegShawty/go_todo_app/pkg"

)

// В фреймворке Gin хэндлер - ф-ция, которая должна иметь в качестве параметра
// указатель на объект gin.Context

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body todo.User true "account info"
// @Success 200 {integer} 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input);err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

// handler.signUp
// h.services -> *service.Service
// h.services.Authorization -> *service.AuthService
// То есть первым CreateUser будет AuthService.CreateUser() - интерфейс Authorization структуры Service 
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input);err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
 
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}