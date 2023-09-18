package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/target"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"fmt"
	"log"

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

	responseUser, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	// buat logic lagi
	// link, err := helper.UploadImage(c)
	// if err != nil {
	// 	log.Printf("Error link: %s", err.Error())
	// 	return helper.FailedRequest(c, err.Error(), nil)
	// }

	// Gabungkan ID pengguna dengan data target
	// newTarget.Proofs = link

	//CEK devisi
	newTarget.DevisiID = responseUser.DevisiID

	//user id penerima -> dari param yang dikasi fe jadi dari node js
	idParam := c.Param("user_id")
	newTarget.UserIDPenerima = idParam

	//mappingg dari request to EntityTarget
	input := TargetRequestToEntity(newTarget)
	input.UserIDPembuat = userID
	fmt.Println(userID)

	result, err := h.targetService.Create(input)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return helper.InternalError(c, "error insert data", err.Error())
	}

	// Mapping create target to Target Response
	resultResponse := EntityToResponse(result)
	// Kirim respon JSON
	log.Println("Target created successfully")
	return helper.SuccessCreate(c, "success create target", resultResponse)
}

func (h *targetHandler) GetAllTarget(c echo.Context) error {
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID, _, _ := middlewares.ExtractToken(c)
	result, err := usernodejs.GetByIdUser(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	// var targetsResponse []TargetResponse
	// for _, v := range result {
	// 	targetsResponse = append(targetsResponse, EntityToResponse(v))
	// }
	return helper.Found(c, "success create target", result)
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
