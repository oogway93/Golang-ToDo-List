package handler_todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo_list/structs"
)

type getAllListsResponse struct {
	Data []structs.ToDo `json:"data"`
}

// @Summary Get All ToDo List
// @Security ApiKeyAuth
// @Tags list
// @Description Getting ALl ToDo List
// @ID get-all-list
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list [get]
func (h *Handler) getList(c *gin.Context) {
	//userId, err := GetUserId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	lists, err := h.services.ToDoList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Create ToDo List
// @Security ApiKeyAuth
// @Tags list
// @Description Create ToDo List
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body structs.ToDoAdd true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list [post]
func (h *Handler) createList(c *gin.Context) {
	var itemList structs.ToDo

	if err := c.BindJSON(&itemList); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.ToDoList.Create(itemList)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := successMessageResponse()
	c.JSON(http.StatusOK, output)
}

// @Summary Get ToDo List
// @Security ApiKeyAuth
// @Tags list
// @Description Get Certain ToDo List By Id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "ToDo List Id"
// @Success 200 {object} structs.ListToDoItems
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/{id} [get]
func (h *Handler) getListByID(c *gin.Context) {
	itemId, _ := strconv.Atoi(c.Param("id"))

	list, err := h.services.GetById(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Update ToDo List
// @Security ApiKeyAuth
// @Tags list
// @Description Update ToDo List By Id
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id path int true "ToDo List Id"
// @Param input body structs.UpdateToDo true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/{id} [put]
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

// @Summary Delete ToDo List
// @Security ApiKeyAuth
// @Tags list
// @Description Delete Certain ToDo List By Id
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param id path int true "ToDo List Id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/{id} [delete]
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
