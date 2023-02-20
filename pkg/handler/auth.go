package handler

import (
	"net/http"

	"github.com/fm2901/go-todo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
