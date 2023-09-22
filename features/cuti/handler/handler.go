package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/cuti"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"strconv"
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
func (handler *CutiHandler) Edit(c echo.Context) error {

	id_cuti := c.Param("id_cuti")
	idUser, _, _ := middlewares.ExtractToken(c)

	var request CutiRequest
	errBind := c.Bind(&request)
	if errBind != nil {
		return helper.FailedRequest(c, "error binding data", nil)
	}

	_, errFile := c.FormFile("image")
	var link string
	var errLink error
	if errFile == nil {
		link, errLink = helper.UploadImage(c)
		if errLink != nil {
			return helper.FailedRequest(c, errLink.Error(), nil)
		}
	}
	entity := RequestToEntity(request)
	entity.UrlPendukung = link
	err := handler.cutiHandler.Edit(entity, id_cuti, idUser)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.FailedRequest(c, err.Error(), nil)
		} else {
			return helper.InternalError(c, err.Error(), nil)
		}
	}
	return helper.SuccessWithOutData(c, "success update data cuti")
}

func (handler *CutiHandler) GetAll(c echo.Context) error {
	var qparams cuti.QueryParams
	page := c.QueryParam("page")
	itemsPerPage := c.QueryParam("itemsPerPage")

	if itemsPerPage == "" {
		qparams.IsClassDashboard = false
	} else {
		qparams.IsClassDashboard = true
		itemsConv, errItem := strconv.Atoi(itemsPerPage)
		if errItem != nil {
			return helper.FailedRequest(c, "item per page not valid", nil)
		}
		qparams.ItemsPerPage = itemsConv
	}
	if page == "" {
		qparams.Page = 1
	} else {
		pageConv, errPage := strconv.Atoi(page)
		if errPage != nil {
			return helper.FailedRequest(c, "page not valid", nil)
		}
		qparams.Page = pageConv
	}

	searchName := c.QueryParam("searchName")
	qparams.SearchName = searchName
	idUser, _, _ := middlewares.ExtractToken(c)

	token,errToken:=usernodejs.GetTokenHandler(c)
	if errToken != nil{
		return helper.Forbidden(c,"token tidak ditemukan",nil)
	}
	bol, data, err := handler.cutiHandler.Get(token,idUser,qparams)

	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}
	var response []CutiResponse
	for _, value := range data {
		response = append(response, EntityToResponse(value))
	}
	return helper.SuccessGetAll(c, "get all cuti successfully", response, bol)
}

func (handler *CutiHandler) GetById(c echo.Context) error {
	id := c.Param("id_cuti")
	data, err := handler.cutiHandler.GetCutiById(id)
	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}
	response := EntityToResponse(data)
	return helper.Success(c, "success get by id cuti", response)
}

func (handler *CutiHandler) Delete(c echo.Context) error {
	id := c.Param("id_cuti")
	err := handler.cutiHandler.Delete(id)
	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}
	return helper.SuccessWithOutData(c, "success delete cuti")
}

func New(handler cuti.CutiServiceInterface) *CutiHandler {
	return &CutiHandler{
		cutiHandler: handler,
	}
}
