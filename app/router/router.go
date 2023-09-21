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

	apinodejs "be_golang/klp3/features/apiNodejs"
	_targetRepo "be_golang/klp3/features/target/data"
	_targetHandler "be_golang/klp3/features/target/handler"
	_targetService "be_golang/klp3/features/target/service"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo, db *gorm.DB, redis *redis.Client) {
	externalAPI := apinodejs.NewExternalData("https://project2.otixx.online")
	dataRes := dataR.New(db,redis)
	serviceRes := serviceR.New(dataRes)
	handlerRes := handlerR.New(serviceRes)

	c.POST("/reimbursements", handlerRes.Add, middlewares.JWTMiddleware())
	c.PUT("/reimbursements/:id_reimbursement", handlerRes.Edit, middlewares.JWTMiddleware())
	c.GET("/reimbursements", handlerRes.GetAll, middlewares.JWTMiddleware())
	c.DELETE("/reimbursements/:id_reimbursement", handlerRes.Delete, middlewares.JWTMiddleware())
	c.GET("/reimbursements/:id_reimbursement", handlerRes.GetById, middlewares.JWTMiddleware())

	dataCuti := dataC.New(db)
	serviceCuti := serviceC.New(dataCuti)
	handlerCuti := handlerC.New(serviceCuti)

	c.POST("/cutis", handlerCuti.AddCuti, middlewares.JWTMiddleware())
	c.GET("/cutis", handlerCuti.GetAll, middlewares.JWTMiddleware())
	c.PUT("/cutis/:id_cuti", handlerCuti.Edit, middlewares.JWTMiddleware())

	dataAbsensi := dataA.New(db, externalAPI)
	serviceAbsensi := serviceA.New(dataAbsensi)
	handlerAbsensi := handlerA.New(serviceAbsensi)

	c.POST("/absensis", handlerAbsensi.Add, middlewares.JWTMiddleware())
	c.PUT("/absensis/:id_absensi", handlerAbsensi.Edit, middlewares.JWTMiddleware())
	c.GET("/absensis", handlerAbsensi.GetAllAbsensi, middlewares.JWTMiddleware())
	c.GET("/absensis/:id_absensi", handlerAbsensi.GetAbsensiById, middlewares.JWTMiddleware())

	targetRepo := _targetRepo.New(db, externalAPI)
	targetService := _targetService.New(targetRepo)
	targetHandlerAPI := _targetHandler.New(targetService)

	c.POST("/user/:user_id/targets", targetHandlerAPI.CreateTarget, middlewares.JWTMiddleware())
	c.GET("/targets", targetHandlerAPI.GetAllTarget, middlewares.JWTMiddleware())
	c.GET("/targets/:target_id", targetHandlerAPI.GetTargetById, middlewares.JWTMiddleware())
	c.PUT("/targets/:target_id", targetHandlerAPI.UpdateTargetById, middlewares.JWTMiddleware())
	c.DELETE("/targets/:target_id", targetHandlerAPI.DeleteTargetById, middlewares.JWTMiddleware())
}
