package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/reimbusment"
	"be_golang/klp3/helper"
	"fmt"
	"strconv"
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

func (handler *ReimbusmentHandler)GetAll(c echo.Context)error{
	var qparams reimbusment.QueryParams
	page:=c.QueryParam("page")
	itemsPerPage:=c.QueryParam("itemsPerPage")

	if itemsPerPage ==""{
		qparams.IsClassDashboard=false
	}else{
		qparams.IsClassDashboard=true
		itemsConv,errItem:=strconv.Atoi(itemsPerPage)
		if errItem != nil{
			return helper.FailedRequest(c,"item per page not valid",nil)
		}
		qparams.ItemsPerPage=itemsConv		
	}
	if page==""{
		qparams.Page=1
	}else{
		pageConv,errPage:=strconv.Atoi(page)
		if errPage != nil{
			return helper.FailedRequest(c,"page not valid",nil)
		}
		qparams.Page=pageConv
	}
	
	searchName:=c.QueryParam("searchName")
	qparams.SearchName=searchName
	idUser,_,_:=middlewares.ExtractToken(c)
	bol,data,err:=handler.reimbushmentHandler.Get(idUser,qparams)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	var response []ReimbursementResponse
	for _,value:=range data{
		response = append(response, EntityToResponse(value))
	}
	return helper.SuccessGetAll(c,"get all reimbursement successfully",response,bol)
}

func (handler *ReimbusmentHandler)Delete(c echo.Context)error{
	id:=c.Param("id_reimbursement")
	err:=handler.reimbushmentHandler.Delete(id)
	if err != nil{
		return helper.InternalError(c,err.Error(),nil)
	}
	return helper.SuccessWithOutData(c,"success delete reimbursement")
}

func New(handler reimbusment.ReimbusmentServiceInterface) *ReimbusmentHandler {
	return &ReimbusmentHandler{
		reimbushmentHandler: handler,
	}
}
