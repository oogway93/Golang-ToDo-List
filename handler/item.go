package handler_todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo_list/structs"
)

// @Summary Get ToDo Item Of List
// @Tags item
// @Description Get ToDo Item Of List
// @ID get-item
// @Accept  json
// @Produce  json
// @Param id path int true "ToDo Item Of List Id"
// @Success 200 {object} []structs.ToDoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/{id}/item [get]
// @Security BearerAuth
func (h *Handler) getItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.ToDoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

// @Summary Create ToDo Item of List
// @Tags item
// @Description Create ToDo Item of List
// @ID create-item
// @Accept  json
// @Produce  json
// @Param id path int true "ToDo Item Of List Id"
// @Param input body structs.ToDoItemAdd true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/{id}/item [post]
// @Security BearerAuth
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
