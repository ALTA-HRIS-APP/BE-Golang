package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/cuti"
	"be_golang/klp3/helper"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type CutiHandler struct {
	cutiHandler cuti.CutiServiceInterface
}

func (handler *CutiHandler) AddCuti(c echo.Context) error {
	idUser, _, _ := middlewares.ExtractToken(c)
	link, errUpload := helper.UploadImage(c)
	if errUpload != nil {
		return helper.FailedRequest(c, errUpload.Error(), nil)
	}

	var input CutiRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return helper.FailedNotFound(c, "error binding", nil)
	}
	entity := RequestToEntity(input)
	entity.UserID = idUser
	entity.UrlPendukung = link
	fmt.Println("cuti", entity)
	err := handler.cutiHandler.Add(entity)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.FailedRequest(c, err.Error(), nil)
		} else {
			return helper.InternalError(c, err.Error(), nil)
		}
	}
	return helper.SuccessWithOutData(c, "success create cuti")
}

func New(handler cuti.CutiServiceInterface) *CutiHandler {
	return &CutiHandler{
		cutiHandler: handler,
	}
}
