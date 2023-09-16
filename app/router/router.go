package router

import (
	"be_golang/klp3/app/middlewares"
	dataC "be_golang/klp3/features/cuti/data"
	handlerC "be_golang/klp3/features/cuti/handler"
	serviceC "be_golang/klp3/features/cuti/service"
	dataR "be_golang/klp3/features/reimbusment/data"
	handlerR "be_golang/klp3/features/reimbusment/handler"
	serviceR "be_golang/klp3/features/reimbusment/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo, db *gorm.DB) {
	dataRes := dataR.New(db)
	serviceRes := serviceR.New(dataRes)
	handlerRes := handlerR.New(serviceRes)

	c.POST("/reimbursments", handlerRes.Add, middlewares.JWTMiddleware())
	c.PUT("/reimbursments/:id_reimbusherment", handlerRes.Edit, middlewares.JWTMiddleware())

	dataCuti := dataC.New(db)
	serviceCuti := serviceC.New(dataCuti)
	handlerCuti := handlerC.New(serviceCuti)

	c.POST("/cutis", handlerCuti.AddCuti, middlewares.JWTMiddleware())
}
