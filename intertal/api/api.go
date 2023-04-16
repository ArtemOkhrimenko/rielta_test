package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate mockgen -source=api.go -destination mock.app.api_test.go -package api_test

type Application interface {
	GetNumber() int64
	PutNumber(num int64) int64
}

type api struct {
	app Application
}

func New(app Application) http.Handler {
	router := gin.New()

	api := api{
		app: app,
	}

	numberHandler := router.Group("/number")

	{
		numberHandler.GET("/", api.getNumber)
		numberHandler.PUT("/", api.putNumber)
	}

	return router
}
