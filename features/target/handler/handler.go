package handler

import (
	"be_golang/klp3/app/middleware"
	"be_golang/klp3/helper"

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
	NewTarget := new(TargetRequest)

	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	NewTarget.UserIDPembuat = middleware.ExtractToken(c)

	//mendapatkan data yang dikirim oleh FE melalui request
	err := c.Bind(&NewTarget)
	if err != nil {
		return helper.FailedRequest(c, "error bind data", nil)
	}

	//mappingg dari request to EntityTarget
	input := TargetRequestToEntity(*NewTarget)

	_, err = h.targetService.Create(input)
	if err != nil {
		return helper.InternalError(c, "error insert data", nil)
	}
	return helper.SuccessCreate(c, "success create target", nil)
}
