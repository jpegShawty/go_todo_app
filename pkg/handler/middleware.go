package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authotizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	// получим header авторизации
	header := c.GetHeader(authotizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		// StatusCode = 401 - Unauthorized
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//parse token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil{
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}