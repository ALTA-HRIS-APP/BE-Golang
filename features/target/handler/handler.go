package handler

import (
	"be_golang/klp3/app/middleware"
	"be_golang/klp3/features/target"
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
	newTarget := new(TargetRequest)

	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID := middleware.ExtractToken(c)
	log.Printf("UserID: %s", userID)

	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&newTarget)
	if err != nil {
		log.Printf("Error binding data: %s", err.Error())
		return helper.FailedRequest(c, "error bind data", nil)
	}

	// Gabungkan ID pengguna dengan data target
	newTarget.UserIDPembuat = userID.ID

	//mappingg dari request to EntityTarget
	input := TargetRequestToEntity(*newTarget)

	err = h.targetService.Create(input)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return helper.InternalError(c, "error insert data", nil)
	}
	log.Println("Target created successfully")
	return helper.SuccessCreate(c, "success create target", nil)
}
