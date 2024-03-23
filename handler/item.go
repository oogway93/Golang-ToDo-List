package handler_todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo_list/structs"
)

type getAllItemsResponse struct {
	Data []structs.ToDoItem `json:"data"`
}

func (h *Handler) getItem(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.ToDoItem.GetAll(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) createItem(c *gin.Context) {}
