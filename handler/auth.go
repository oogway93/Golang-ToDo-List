package handler_todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/structs"
)

// @Summary Sign Up User
// @Tags auth
// @Description Sign Up User
// @ID signUp-user
// @Accept json
// @Produce json
// @Param input body structs.SignUpUser true "user info"
// @Success 200 {object} statusAuthResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user structs.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.services.User.CreateUser(user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output := authMessageResponse()
	c.JSON(http.StatusOK, output)
}

// @Summary Sign In User
// @Tags auth
// @Description Sign In User
// @ID signIn-user
// @Accept json
// @Produce json
// @Param input body structs.SignInUser true "user info"
// @Success 200 {object} statusAuthResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var user structs.SignInUser
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.services.User.GenerateToken(user.Username, user.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
