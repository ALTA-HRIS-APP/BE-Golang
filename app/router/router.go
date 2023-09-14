package router

import (
	dataR "be_golang/klp3/features/reimbusment/data"
	handlerR "be_golang/klp3/features/reimbusment/handler"
	serviceR "be_golang/klp3/features/reimbusment/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(c *echo.Echo,db *gorm.DB){
	dataRes:=dataR.New(db)
	serviceRes:=serviceR.New(dataRes)
	handlerRes:=handlerR.New(serviceRes)

	c.POST("/reimbursments",handlerRes.Add)
	c.PUT("/reimbursments/:id_reimbusherment",handlerRes.Edit)
}