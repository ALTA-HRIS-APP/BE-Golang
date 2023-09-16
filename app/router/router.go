package router

import (
	"be_golang/klp3/app/middleware"
	dataR "be_golang/klp3/features/reimbusment/data"
	handlerR "be_golang/klp3/features/reimbusment/handler"
	serviceR "be_golang/klp3/features/reimbusment/service"

	_targetRepo "be_golang/klp3/features/target/data"
	_targetHandler "be_golang/klp3/features/target/handler"
	_targetService "be_golang/klp3/features/target/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo, db *gorm.DB) {
	dataRes := dataR.New(db)
	serviceRes := serviceR.New(dataRes)
	handlerRes := handlerR.New(serviceRes)

	c.POST("/reimbursments", handlerRes.Add)
	c.PUT("/reimbursments/:id_reimbusherment", handlerRes.Edit)

	targetRepo := _targetRepo.New(db)
	targetService := _targetService.New(targetRepo)
	targetHandlerAPI := _targetHandler.New(targetService)
	c.POST("/targets", targetHandlerAPI.CreateTarget, middleware.JWTMiddleware())
	// e.GET("/targets", targetHandlerAPI.GetAllTarget, middlewares.JWTMiddleware())
	// e.GET("/targets/:target_id", targetHandlerAPI.GetTargetById, middlewares.JWTMiddleware())
	// e.PUT("/targets/:target_id", targetHandlerAPI.UpdateTargetById, middlewares.JWTMiddleware())
	// e.DELETE("/targets/:target_id", targetHandlerAPI.DeleteTargetById, middlewares.JWTMiddleware())
}
