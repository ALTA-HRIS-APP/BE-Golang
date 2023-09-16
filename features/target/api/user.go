package api

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

func ApiGetUser(token string) (ResponseUser, error) {
	// Create a Resty Client
	client := resty.New()
	response := ResponseUser{}

	// POST JSON string
	// No need to set content type, if you have client level setting
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		// SetBody(`{"email": "admin@gmail.com", "password":"12345678"}`).
		SetHeader("Authorization", "Bearer "+token).
		// SetQueryParams(params).
		SetError(&response).
		SetResult(&response). // or SetResult(AuthSuccess{}).
		Get("http://pintu2.otixx.online/user/profile/")
	if err != nil {
		print(err)
		return ResponseUser{}, errors.New(response.Meta.Message)
	}
	return response, nil
}
func ApiGetUsers(token string) (ResponseUsers, error) {
	// Create a Resty Client
	client := resty.New()
	// var response any
	response := ResponseUsers{}

	// POST JSON string
	// No need to set content type, if you have client level setting
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		// SetBody(`{"email": "admin@gmail.com", "password":"12345678"}`).
		SetHeader("Authorization", "Bearer "+token).
		// SetQueryParams(params).
		SetError(&response).
		SetResult(&response). // or SetResult(AuthSuccess{}).
		Get("http://pintu2.otixx.online/user/")
	if err != nil {
		print(err)
		// return ResponseUsers{}, errors.New(response.Meta.Message)
	}
	return response, nil
}

type ResponseUser struct {
	Meta Meta `json:"meta"`
	Data User `json:"data,omitempty"`
}
type ResponseUsers struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    []struct {
			ID          string    `json:"id"`
			NamaLengkap string    `json:"nama_lengkap"`
			Surel       string    `json:"surel"`
			NoHp        string    `json:"no_hp"`
			Jabatan     string    `json:"jabatan"`
			KataSandi   string    `json:"kata_sandi"`
			Status      bool      `json:"status"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			DevisiID    string    `json:"devisiId"`
			RoleID      string    `json:"roleId"`
			Devisi      struct {
				ID        string    `json:"id"`
				Nama      string    `json:"nama"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"devisi"`
			Role struct {
				ID        string    `json:"id"`
				Nama      string    `json:"nama"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"role"`
		} `json:"data"`
	} `json:"meta"`
}
type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type User struct {
	ID          string    `json:"id"`
	NamaLengkap string    `json:"nama_lengkap"`
	Surel       string    `json:"surel"`
	NoHp        string    `json:"no_hp"`
	Jabatan     string    `json:"jabatan"`
	KataSandi   string    `json:"kata_sandi"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DevisiID    string    `json:"devisiId"`
	RoleID      string    `json:"roleId"`
	Devisi      struct {
		ID        string    `json:"id"`
		Nama      string    `json:"nama"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"devisi"`
	Role struct {
		ID        string    `json:"id"`
		Nama      string    `json:"nama"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"role"`
}
