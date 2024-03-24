package handler_todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/structs"
)

// @Summary Create User
// @Tags auth
// @Description Create User
// @ID create-user
// @Accept json
// @Produce json
// @Param input body structs.UserAdd true "user info"
// @Success 200 {object} statusAuthResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user structs.UserAdd

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.services.SignUp(user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := authMessageResponse()
	c.JSON(http.StatusOK, output)
}
