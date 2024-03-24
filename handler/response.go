package handler_todo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type statusAuthResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func successMessageResponse() *statusResponse {
	response := &statusResponse{"success"}
	return response
}

func authMessageResponse() *statusAuthResponse {
	response := &statusAuthResponse{"Authentication is success"}
	return response
}
