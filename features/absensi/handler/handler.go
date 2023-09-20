package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/absensi"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AbsensiHandler struct {
	absensiService absensi.AbsensiServiceInterface
}

func New(service absensi.AbsensiServiceInterface) *AbsensiHandler {
	return &AbsensiHandler{
		absensiService: service, // Mengganti absensiService dengan service
	}
}

func (handler *AbsensiHandler) Edit(c echo.Context) error {
	// idUser,_,_:=middlewares.ExtractToken(c)
	idUser := "13947f80-78b9-446f-9fe4-cb25caa4bea4"
	id := c.Param("id_absensi")
	err := handler.absensiService.Edit(idUser, id)
	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}

	return helper.SuccessWithOutData(c, "success update absen pulang")
}

func (handler *AbsensiHandler) Add(c echo.Context) error {
	// idUser,_,_:=middlewares.ExtractToken(c)
	idUser := "13947f80-78b9-446f-9fe4-cb25caa4bea4"
	err := handler.absensiService.Add(idUser)
	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}
	return helper.SuccessCreate(c, "success create absen", nil)
}

func (handler *AbsensiHandler) GetAllAbsensi(c echo.Context) error {
	var qparams absensi.QueryParams
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
	bol, data, err := handler.absensiService.Get(idUser, qparams)
	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}
	var response []AbsensiResponse
	for _, value := range data {
		response = append(response, EntityToResponse(value))

	}
	return helper.SuccessGetAll(c, "get all reimbursement successfully", response, bol)
}

func (handler *AbsensiHandler) GetAbsensiById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)
	apiUser, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	idParam := c.Param("absensi_id")
	result, err := handler.absensiService.GetById(idParam, apiUser.ID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	resultResponse := EntityToResponse(result)
	return helper.Success(c, "success create absensi", resultResponse)
}
