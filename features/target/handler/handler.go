package handler

import (
	"be_golang/klp3/app/middlewares"
	"be_golang/klp3/features/target"
	"be_golang/klp3/helper"
	"log"
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
	// Mengambil ID pengguna dari token JWT yang terkait dengan permintaan
	userID, _, _ := middlewares.ExtractToken(c)
	// helper.PrettyPrint(user)
	log.Println("UserID: ", userID)

	//mengecek user id dari get by id user id api node js
	userProfile, err := h.targetService.GetUserByIDAPI(userID)
	if err != nil {
		log.Printf("Error get detail user: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	newTarget := TargetRequest{}
	//mendapatkan data yang dikirim oleh FE melalui request
	err = c.Bind(&newTarget)
	if err != nil {
		log.Printf("Error binding data: %s", err.Error())
		return helper.FailedRequest(c, "Failed to bind data", nil)
	}

	//mengisi user id pembuat dengan user id ang login
	newTarget.UserIDPembuat = userID

	//mengisi divisi id dengan divisi user yang login
	newTarget.DevisiID = userProfile.DevisiID

	//user id penerima -> dari param yang dikasi fe jadi dari node js
	idParam := c.Param("user_id")
	newTarget.UserIDPenerima = idParam

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
			// Handle the error, if any
			return helper.FailedRequest(c, errLink.Error(), nil)
		}
		newTarget.Proofs = cloudnaryLink
	}
	//mappingg dari request to EntityTarget
	targetInput := TargetRequestToEntity(newTarget)
	targetID, err := h.targetService.Create(targetInput)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return helper.InternalError(c, "Failed to insert data", err.Error())
	}
	targetInput.ID = targetID
	// Mapping create target to Target Response
	responseTarget := EntityToResponse(targetInput)
	// Kirim respon JSON
	log.Println("Target created successfully")
	return helper.SuccessCreate(c, "success create target", responseTarget)
}

func (h *targetHandler) GetAllTarget(c echo.Context) error {
	// Get user ID from the JWT token associated with the request
	var qParam target.QueryParam
	page := c.QueryParam("page")
	limitPerPage := c.QueryParam("limitPerPage")

	if limitPerPage != "" {
		qParam.ExistOtherPage = true
		limitConv, err := strconv.Atoi(limitPerPage)
		if err != nil {
			log.Printf("Invalid limit item per page: %s", err.Error())
			return helper.FailedRequest(c, "Invalid limit item per page", nil)
		}
		qParam.LimitPerPage = limitConv
	}
	if page != "" {
		pageConv, err := strconv.Atoi(page)
		if err != nil {
			log.Printf("Invalid page: %s", err.Error())
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
		log.Printf("Internal server error: %s", err.Error())
		return helper.InternalError(c, err.Error(), nil)
	}

	var targetsResponse []TargetResponse
	for _, v := range data {
		targetsResponse = append(targetsResponse, EntityToResponse(v))
	}

	log.Println("Get all targets successfully")
	return helper.Success(c, "Get all targets successfully", map[string]interface{}{
		"targets":  targetsResponse,
		"nextPage": nextPage,
	})
}

// GetTargetById retrieves target details by its ID.
func (h *targetHandler) GetTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)

	// Get user details
	_, err := h.targetService.GetUserByIDAPI(userID)
	if err != nil {
		log.Printf("Error getting user details: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	idParam := c.Param("target_id")

	// Get target details and check if the user has permission to access it
	result, err := h.targetService.GetById(idParam, userID)
	if err != nil {
		log.Printf("Error getting target details: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	resultResponse := EntityToResponse(result)
	log.Println("Get target by ID successfully")
	return helper.Found(c, "Success getting target details", resultResponse)
}

func (h *targetHandler) UpdateTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)
	apiUser, err := h.targetService.GetUserByIDAPI(userID)
	if err != nil {
		log.Printf("Error getting user details: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	idParam := c.Param("target_id")
	_, err = h.targetService.GetById(idParam, apiUser.ID)
	if err != nil {
		log.Printf("Error getting target details: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	// Get input data from the user
	inputTarget := TargetRequest{}
	errBind := c.Bind(&inputTarget)
	if errBind != nil {
		return helper.FailedRequest(c, "failed to bind target data", err.Error())
	}
	//mengisi proof dengan link dari cloudnary
	_, errFile := c.FormFile("image")
	if errFile != nil {
		// Handle the error, if any
		return helper.FailedRequest(c, errFile.Error(), nil)
	}

	cloudnaryLink, errLink := helper.UploadImage(c)
	if errLink != nil {
		// Handle the error, if any
		return helper.FailedRequest(c, errLink.Error(), nil)
	}
	inputTarget.Proofs = cloudnaryLink
	// Map target request to entity
	entityTarget := TargetRequestToEntity(inputTarget)

	// Perform the target data update in the service
	err = h.targetService.UpdateById(idParam, userID, entityTarget)
	if err != nil {
		log.Printf("Error updating target: %s", err.Error())
		return helper.InternalError(c, "failed to update target data", err.Error())
	}
	// Get the updated target data for response
	updatedTarget, err := h.targetService.GetById(idParam, userID)
	if err != nil {
		log.Printf("Error getting updated target details: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}
	// Map the updated target to Target Response
	resultResponse := EntityToResponse(updatedTarget)
	// Send JSON response
	return helper.Success(c, "target updated successfully", resultResponse)
}

// DeleteTargetById handles the deletion of a target by its ID.
func (h *targetHandler) DeleteTargetById(c echo.Context) error {
	userID, _, _ := middlewares.ExtractToken(c)

	// Get user details
	_, err := h.targetService.GetUserByIDAPI(userID)
	if err != nil {
		log.Printf("Error getting user details: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	idParam := c.Param("target_id")

	// Check if the target exists and is allowed to be deleted
	err = h.targetService.DeleteById(idParam, userID)
	if err != nil {
		log.Printf("Error deleting target: %s", err.Error())
		return helper.FailedRequest(c, err.Error(), nil)
	}

	log.Println("Target deleted successfully")
	return helper.Success(c, "Target deleted successfully", nil)
}
