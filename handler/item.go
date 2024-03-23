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

func (h *Handler) createItem(c *gin.Context) {
	var item structs.ToDoItem

	if err := c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, _ := strconv.Atoi(c.Param("id"))

	err := h.services.ToDoItem.Create(listId, item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := successMessageResponse()
	c.JSON(http.StatusOK, output)
}
