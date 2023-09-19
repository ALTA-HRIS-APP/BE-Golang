package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/target"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type targetHandler struct {
	targetService target.TargetServiceInterface
}

func New(service target.TargetServiceInterface) *targetHandler {
	return &targetHandler{
		targetService: service,
	}
}

func (h *targetHandler) CreateTarget(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID, _, _ := middlewares.ExtractToken(c)
	// helper.PrettyPrint(user)
	log.Println("UserID: ", userID)

	newTarget := TargetRequest{}
	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&newTarget)
	if err != nil {
		log.Printf("Error binding data: %s", err.Error())
		return helper.FailedRequest(c, "error bind data", nil)
	}
	//mengecek user id dari get by id user id api node js
	responseUser, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	//mengisi user id pembuat dengan user id ang login
	newTarget.UserIDPembuat = userID

	//mengisi divisi id dengan divisi user yang login
	newTarget.DevisiID = responseUser.DevisiID

	//user id penerima -> dari param yang dikasi fe jadi dari node js
	idParam := c.Param("user_id")
	newTarget.UserIDPenerima = idParam

	//mengisi proof dengan link dari cloudnary
	if newTarget.Proofs != "" {
		link, err := helper.UploadImage(c)
		if err != nil {
			log.Printf("Error link: %s", err.Error())
			return helper.FailedRequest(c, err.Error(), nil)
		}
		newTarget.Proofs = link
	}

	//mappingg dari request to EntityTarget
	input := TargetRequestToEntity(newTarget)
	targetID, err := h.targetService.Create(input)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return helper.InternalError(c, "error insert data", err.Error())
	}
	input.ID = targetID
	// Mapping create target to Target Response
	resultResponse := EntityToResponse(input)
	// Kirim respon JSON
	log.Println("Target created successfully")
	return helper.SuccessCreate(c, "success create target", resultResponse)
}

func (h *targetHandler) GetAllTarget(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	var qParam target.QueryParam
	page := c.QueryParam("page")
	limitPerPage := c.QueryParam("limitPerPage")

	if limitPerPage != "" {
		qParam.ExistOtherPage = true
		limitConv, err := strconv.Atoi(limitPerPage)
		if err != nil {
			return helper.FailedRequest(c, "limit item per page not valid", nil)
		}
		qParam.LimitPerPage = limitConv
	}
	if page != "" {
		pageConv, err := strconv.Atoi(page)
		if err != nil {
			return helper.FailedRequest(c, "page not valid", nil)
		}
		qParam.Page = pageConv
	} else {
		qParam.Page = 1
	}
	searchKonten := c.QueryParam("searchKonten")
	qParam.SearchKonten = searchKonten

	searchStatus := c.QueryParam("searchStatus")
	qParam.SearchStatus = searchStatus

	userID, _, _ := middlewares.ExtractToken(c)
	count, data, nextPage, err := h.targetService.GetAll(userID, qParam)

	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}

	var targetsResponse []TargetResponse
	for _, v := range data {
		targetsResponse = append(targetsResponse, EntityToResponse(v))
	}
	response := map[string]interface{}{
		"totalData": count,
		"data":      targetsResponse,
		"nextPage":  nextPage,
	}
	return helper.Success(c, "get all target successfully", response)
}

func (h *targetHandler) GetTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)
	apiUser, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	idParam := c.Param("target_id")
	result, err := h.targetService.GetById(idParam, apiUser.ID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	resultResponse := EntityToResponse(result)
	return helper.Found(c, "success create target", resultResponse)
}
func (h *targetHandler) UpdateTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)
	apiUser, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	idParam := c.Param("target_id")
	_, err = h.targetService.GetById(idParam, apiUser.ID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	//mengambil data input dari user id penerima
	inputTarget := TargetReqPenerima{}
	errBind := c.Bind(&inputTarget)
	if errBind != nil {
		return helper.FailedRequest(c, "success create target", err.Error())
	}
	//Mapping targeet request to entity
	entityTarget := TargetReqPenerimaToEntity(inputTarget)

	// Melakukan pembaruan data target di service
	err = h.targetService.UpdateById(idParam, apiUser.ID, entityTarget)
	if err != nil {
		log.Printf("Error update target: %s", err.Error())
		return helper.InternalError(c, "error updated data", err.Error())
	}
	// Mendapatkan data proyek yang telah diperbarui untuk respon
	updatedTarget, err := h.targetService.GetById(idParam, userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	// Mapping updated target to Target Response
	resultResponse := EntityToResponse(updatedTarget)
	// Kirim respon JSON
	return helper.Success(c, "target updated successfully", resultResponse)
}
func (h *targetHandler) DeleteTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)
	apiUser, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	idParam := c.Param("target_id")
	_, err = h.targetService.GetById(idParam, apiUser.ID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	return helper.Success(c, "target deleted successfully", nil)
}
