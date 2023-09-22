package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/target"
	"be_golang/klp3/helper"
	"net/http"
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
	userID, _, _ := middlewares.ExtractToken(c)

	newTarget := TargetRequest{}
	err := c.Bind(&newTarget)
	if err != nil {
		return helper.FailedRequest(c, "Failed to bind data", nil)
	}

	newTarget.UserIDPembuat = userID

	idParam := c.Param("user_id")
	newTarget.UserIDPenerima = idParam

	_, errFile := c.FormFile("image")
	if errFile != nil && errFile != http.ErrMissingFile {
		return helper.FailedRequest(c, errFile.Error(), nil)
	}

	if errFile == nil {
		// Jika ada file gambar yang diunggah, maka unggah ke Cloudinary
		cloudnaryLink, errLink := helper.UploadImage(c)
		if errLink != nil {
			return helper.FailedRequest(c, errLink.Error(), nil)
		}
		newTarget.Proofs = cloudnaryLink
	}
	//mappingg dari request to EntityTarget
	targetInput := TargetRequestToEntity(newTarget)
	targetID, err := h.targetService.Create(targetInput)
	if err != nil {
		return helper.InternalError(c, "Failed to insert data", err.Error())
	}
	targetInput.ID = targetID
	// Mapping create target to Target Response
	responseTarget := EntityToResponse(targetInput)
	return helper.SuccessCreate(c, "success create target", responseTarget)
}

func (h *targetHandler) GetAllTarget(c echo.Context) error {
	var qParam target.QueryParam
	page := c.QueryParam("page")
	limitPerPage := c.QueryParam("limitPerPage")

	if limitPerPage != "" {
		qParam.ExistOtherPage = true
		limitConv, err := strconv.Atoi(limitPerPage)
		if err != nil {
			return helper.FailedRequest(c, "Invalid limit item per page", nil)
		}
		qParam.LimitPerPage = limitConv
	}
	if page != "" {
		pageConv, err := strconv.Atoi(page)
		if err != nil {
			return helper.FailedRequest(c, "Invalid page", nil)
		}
		qParam.Page = pageConv
	} else {
		qParam.Page = 1
	}
	searchKonten := c.QueryParam("search_konten")
	qParam.SearchKonten = searchKonten

	searchStatus := c.QueryParam("search_status")
	qParam.SearchStatus = searchStatus

	userID, _, _ := middlewares.ExtractToken(c)
	nextPage, data, err := h.targetService.GetAll(userID, qParam)

	if err != nil {
		return helper.InternalError(c, err.Error(), nil)
	}

	var targetsResponse []TargetResponse
	for _, v := range data {
		targetsResponse = append(targetsResponse, EntityToResponse(v))
	}

	return helper.SuccessGetAll(c, "get all target successfully", targetsResponse, nextPage)
}

func (h *targetHandler) GetTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)

	idParam := c.Param("target_id")

	result, err := h.targetService.GetById(idParam, userID)
	if err != nil {
		return helper.FailedRequest(c, err.Error(), nil)
	}

	resultResponse := EntityToResponse(result)
	return helper.Found(c, "Success getting target details", resultResponse)
}

func (h *targetHandler) UpdateTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)

	idParam := c.Param("target_id")
	_, err := h.targetService.GetById(idParam, userID)
	if err != nil {
		return helper.FailedRequest(c, err.Error(), nil)
	}

	inputTarget := TargetRequest{}
	errBind := c.Bind(&inputTarget)
	if errBind != nil {
		return helper.FailedRequest(c, "failed to bind target data", err.Error())
	}
	//mengisi proof dengan link dari cloudnary
	_, errFile := c.FormFile("image")
	if errFile != nil && errFile != http.ErrMissingFile {
		// Handle the error, except when it's ErrMissingFile (no file uploaded)
		return helper.FailedRequest(c, errFile.Error(), nil)
	}

	if errFile == nil {
		// Jika ada file gambar yang diunggah, maka unggah ke Cloudinary
		cloudnaryLink, errLink := helper.UploadImage(c)
		if errLink != nil {
			return helper.FailedRequest(c, errLink.Error(), nil)
		}
		inputTarget.Proofs = cloudnaryLink
	}
	// Map target request to entity
	entityTarget := TargetRequestToEntity(inputTarget)

	err = h.targetService.UpdateById(idParam, userID, entityTarget)
	if err != nil {
		return helper.InternalError(c, "failed to update target data", err.Error())
	}
	// Get the updated target data for response
	updatedTarget, err := h.targetService.GetById(idParam, userID)
	if err != nil {
		return helper.FailedRequest(c, err.Error(), nil)
	}
	// Map the updated target to Target Response
	resultResponse := EntityToResponse(updatedTarget)
	return helper.Success(c, "target updated successfully", resultResponse)
}

// DeleteTargetById handles the deletion of a target by its ID.
func (h *targetHandler) DeleteTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)

	idParam := c.Param("target_id")

	err := h.targetService.DeleteById(idParam, userID)
	if err != nil {
		return helper.FailedRequest(c, err.Error(), nil)
	}
	return helper.Success(c, "Target deleted successfully", nil)
}
