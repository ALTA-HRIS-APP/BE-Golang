package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/reimbusment"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type ReimbusmentHandler struct {
	reimbushmentHandler reimbusment.ReimbusmentServiceInterface
}

func (handler *ReimbusmentHandler)Edit(c echo.Context)error{
		// idUser,_:=middleware.ExtractToken(c)
		idUser:="651993b3-fac0-4b92-9ea0-8cba4b66f7e5"
		
		var request ReimbursementRequest
		errBind:=c.Bind(&request)
		if errBind != nil{
			return helper.FailedRequest(c, "error bind data"+errBind.Error(), nil)
		}

		id:=c.Param("id_reimbusherment")

		entity:=RequestToEntity(request)

		if entity.UserID == ""{
			if entity.Status !=""{
				return helper.FailedRequest(c,"hanya admin yang dapat mengedit status",nil)
			}
			if entity.Persetujuan !=""{
				return helper.FailedRequest(c,"hanya HR yang dapat mengedit persetujuan",nil)
			}
			entity.UserID=idUser
			err:=handler.reimbushmentHandler.Edit(entity,id)
			if err != nil{
				return helper.InternalError(c,err.Error(),nil)
			}
		}else{
			if entity.Description !=""{
				return helper.FailedRequest(c,"hanya user yang dapat mengedit description",nil)
			}
			if entity.Tipe !=""{
				return helper.FailedRequest(c,"hanya user yang dapat mengedit type",nil)
			}
			if entity.Nominal !=0{
				return helper.FailedRequest(c,"hanya user yang dapat mengedit nominal",nil)
			}
			if entity.UrlBukti !=""{
				return helper.FailedRequest(c,"hanya user yang dapat mengedit bukti transaksi",nil)
			}
			err:=handler.reimbushmentHandler.EditAdmin(entity.Status,entity.UserID,idUser,id)
			if err != nil{
				return helper.InternalError(c,err.Error(),nil)
			}
		}
		return helper.SuccessWithOutData(c,"success update reimbursment")
}

func (handler *ReimbusmentHandler)Add(c echo.Context)error{
	idUser,_,_:=middlewares.ExtractToken(c)
	// email:=middlewares.ExtractTokenID()
	token,errToken:=usernodejs.GetTokenHandler(c)
	if errToken != nil{
		return helper.UnAutorization(c,"gagal get token",nil)
	}
	fmt.Println("token user",token)
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