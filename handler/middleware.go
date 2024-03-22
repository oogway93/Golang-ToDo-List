package handler_todo

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	listID = "id"
)

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(listID)
	if !ok {
		return 0, errors.New("id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("id is of invalid type")
	}

	return idInt, nil
}
