package handler

import (
	"be_golang/klp3/features/reimbusment"
	"be_golang/klp3/helper"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type ReimbusmentHandler struct {
	reimbushmentHandler reimbusment.ReimbusmentServiceInterface
}

func (handler *ReimbusmentHandler)Add(c echo.Context)error{
	// idUser,_:=middleware.ExtractToken(c)
	idUser:="651993b3-fac0-4b92-9ea0-8cba4b66f7e5"
	var request ReimbursementRequest
	errBind:=c.Bind(&request)
	if errBind != nil{
		return helper.FailedRequest(c, "error bind data"+errBind.Error(), nil)
	}
	link,errLink:=helper.UploadImage(c)
	if errLink != nil{
		return helper.FailedRequest(c, errLink.Error(), nil)
	}

	fmt.Println(request)
	entity:=RequestToEntity(request)
	entity.UserID = idUser
	entity.UrlBukti = link
	err:= handler.reimbushmentHandler.Add(entity)
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.FailedRequest(c,err.Error(),nil)
		}else{
			return helper.InternalError(c,err.Error(),nil)
		}
	}
	return helper.SuccessWithOutData(c,"success create reimbursment")
}
func New(handler reimbusment.ReimbusmentServiceInterface)*ReimbusmentHandler{
	return &ReimbusmentHandler{
		reimbushmentHandler: handler,
	}
}