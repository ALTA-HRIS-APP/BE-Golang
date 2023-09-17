package handler

import (
	"be_golang/klp3/app/middleware"
	"be_golang/klp3/features/target"
	"be_golang/klp3/features/target/api"
	"be_golang/klp3/helper"
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
	user := middleware.ExtractToken(c)
	// helper.PrettyPrint(user)
	log.Printf("UserID: %s", user.ID)

	newTarget := TargetRequest{}
	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&newTarget)
	if err != nil {
		log.Printf("Error binding data: %s", err.Error())
		return helper.FailedRequest(c, "error bind data", nil)
	}

	responseUser, err := api.ApiGetUserProfile(user.Token)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	if user.Role=="user"{
		return helper.UnAutorization(c.Echo().NewContext())
	}
	link, err := helper.UploadImage(c)
	if err != nil {
		log.Printf("Error link: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	// Gabungkan ID pengguna dengan data target
	newTarget.UserIDPembuat = responseUser.Data.ID
	newTarget.Proofs = link
	newTarget.DevisiID = responseUser.Data.Devisi.ID

	//mappingg dari request to EntityTarget
	input := TargetRequestToEntity(newTarget)

	err = h.targetService.Create(input)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return helper.InternalError(c, "error insert data", nil)
	}
	log.Println("Target created successfully")
	return helper.SuccessCreate(c, "success create target", nil)
}
