package router

import (
	"be_golang/klp3/app/middlewares"
	dataC "be_golang/klp3/features/cuti/data"
	handlerC "be_golang/klp3/features/cuti/handler"
	serviceC "be_golang/klp3/features/cuti/service"
	dataR "be_golang/klp3/features/reimbusment/data"
	handlerR "be_golang/klp3/features/reimbusment/handler"
	serviceR "be_golang/klp3/features/reimbusment/service"

	dataA "be_golang/klp3/features/absensi/data"
	handlerA "be_golang/klp3/features/absensi/handler"
	serviceA "be_golang/klp3/features/absensi/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo, db *gorm.DB) {
	dataRes := dataR.New(db)
	serviceRes := serviceR.New(dataRes)
	handlerRes := handlerR.New(serviceRes)

	c.POST("/reimbursements", handlerRes.Add, middlewares.JWTMiddleware())
	c.PUT("/reimbursements/:id_reimbursement", handlerRes.Edit, middlewares.JWTMiddleware())
	c.GET("/reimbursements", handlerRes.GetAll, middlewares.JWTMiddleware())
	c.DELETE("/reimbursements/:id_reimbursement", handlerRes.Delete, middlewares.JWTMiddleware())

	dataCuti := dataC.New(db)
	serviceCuti := serviceC.New(dataCuti)
	handlerCuti := handlerC.New(serviceCuti)

	c.POST("/cutis", handlerCuti.AddCuti, middlewares.JWTMiddleware())
	c.GET("/cutis", handlerCuti.GetAll, middlewares.JWTMiddleware())

	dataAbsensi := dataA.New(db)
	serviceAbsensi := serviceA.New(dataAbsensi)
	handlerAbsensi := handlerA.New(serviceAbsensi)

	c.POST("/absensis", handlerAbsensi.Add)
	c.PUT("/absensis/:id_absensi", handlerAbsensi.Edit)
	c.GET("/absensis", handlerAbsensi.GetAllAbsensi)
	c.GET("/absensis/:id_absensi", handlerAbsensi.GetById)
}
