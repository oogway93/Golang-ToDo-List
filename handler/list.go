package handler_todo

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo_list/structs"
)

type getAllListsResponse struct {
	Data []structs.ToDo `json:"data"`
}

func (h *Handler) getList(c *gin.Context) {
	lists, err := h.services.ToDoList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) createList(c *gin.Context) {
	var itemList structs.ToDo

	if err := c.BindJSON(&itemList); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.Create(itemList)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := successMessageResponse()
	c.JSON(http.StatusOK, output)
}

func (h *Handler) getListByID(c *gin.Context) {
	itemId, _ := strconv.Atoi(c.Param("id"))

	list, err := h.services.GetById(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
func (h *Handler) updateList(c *gin.Context) {
	itemId, _ := strconv.Atoi(c.Param("id"))

	var updatedListItem structs.UpdateToDo
	if err := c.BindJSON(&updatedListItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.services.Update(itemId, updatedListItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := successMessageResponse()
	c.JSON(http.StatusOK, output)
}

func (h *Handler) deleteList(c *gin.Context) {
	itemId, _ := strconv.Atoi(c.Param("id"))

	err := h.services.Delete(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := successMessageResponse()
	c.JSON(http.StatusOK, output)
}
