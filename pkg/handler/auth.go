package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/jpegShawty/go_todo_app/pkg"

)

// В фреймворке Gin хэндлер - ф-ция, которая должна иметь в качестве параметра
// указатель на объект gin.Context
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input);err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}