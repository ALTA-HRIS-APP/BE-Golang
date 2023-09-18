package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/reimbusment"
	"be_golang/klp3/helper"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type ReimbusmentHandler struct {
	reimbushmentHandler reimbusment.ReimbusmentServiceInterface
}


func (handler *ReimbusmentHandler) Add(c echo.Context) error {
	idUser, _, _ := middlewares.ExtractToken(c)

	var request ReimbursementRequest
	errBind := c.Bind(&request)
	if errBind != nil {
		return helper.FailedRequest(c, "error bind data"+errBind.Error(), nil)
	}
	link, errLink := helper.UploadImage(c)
	if errLink != nil {
		return helper.FailedRequest(c, errLink.Error(), nil)
	}

	fmt.Println(request)
	entity := RequestToEntity(request)
	entity.UserID = idUser
	entity.UrlBukti = link
	err := handler.reimbushmentHandler.Add(entity)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.FailedRequest(c, err.Error(), nil)
		} else {
			return helper.InternalError(c, err.Error(), nil)
		}
	}
	return helper.SuccessWithOutData(c, "success create reimbursment")
}

func (handler *ReimbusmentHandler) Edit(c echo.Context)error{

	idRemb:=c.Param("id_reimbursement")
	idUser,_,_:=middlewares.ExtractToken(c)

	var request ReimbursementRequest
	errBind:=c.Bind(&request)	
	if errBind != nil{
		return helper.FailedRequest(c,"error binding data",nil)
	}

	_, errFile := c.FormFile("image")
	var link string
	var errLink error
	if errFile == nil {
		link, errLink = helper.UploadImage(c)
		if errLink != nil {
			return helper.FailedRequest(c,errLink.Error(),nil)
		}
	}
	entity:=RequestToEntity(request)
	entity.UrlBukti=link
	err:=handler.reimbushmentHandler.Edit(entity,idRemb,idUser)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.FailedRequest(c, err.Error(), nil)
		} else {
			return helper.InternalError(c, err.Error(), nil)
		}
	}
	return helper.SuccessWithOutData(c,"success update data reimbursement")
}
func New(handler reimbusment.ReimbusmentServiceInterface) *ReimbusmentHandler {
	return &ReimbusmentHandler{
		reimbushmentHandler: handler,
	}
}
