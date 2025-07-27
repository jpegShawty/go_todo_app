package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/jpegShawty/go_todo_app/pkg"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil{
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	items, err :=  h.services.TodoItem.GetAll(userId, listId)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err :=  h.services.TodoItem.GetAll(userId, itemId)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput

	if err := c.BindJSON(&input); err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})

}
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err =  h.services.TodoItem.Delete(userId, itemId)
	if err != nil{
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})

}