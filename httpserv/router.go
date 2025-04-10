package httpserv

import (
	"card-service/infrastructure"
	"card-service/internal/adaptor/handler"
	repository "card-service/internal/adaptor/repo"
	"card-service/internal/core/service"

	libmiddleware "github.com/basputtipong/library/middleware"
	"github.com/gin-gonic/gin"
)

func bindGetCardRoute(app *gin.Engine) {
	repo := repository.NewCardRepo(infrastructure.DB)
	svc := service.NewCardService(repo)
	hdl := handler.NewCardHandler(svc)

	app.GET("/card", libmiddleware.JWTVerify(), hdl.Handle)
}

func bindHelthRoute(app *gin.Engine) {
	app.GET("/health", handler.HealthHandle)
}
