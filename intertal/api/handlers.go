package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *api) getNumber(c *gin.Context) {

	number := a.app.GetNumber()

	c.JSON(http.StatusOK, number)
}

func (a *api) putNumber(c *gin.Context) {
	req := &CreateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	number := a.app.PutNumber(req.Number)

	c.JSON(http.StatusOK, number)
}
