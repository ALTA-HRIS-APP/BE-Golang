package handler

import (
	"be_golang/klp3/features/reimbusment"
	usernodejs "be_golang/klp3/features/userNodejs"
)

type ReimbursementRequest struct {
	ID              string `json:"id" form:"id"`
	Description     string `json:"description" form:"description"`
	Status          string `json:"status" form:"status"`
	BatasanReimburs int 	`json:"batasan_reimburs" form:"batasan_reimburs"`
	Nominal         int 	`json:"nominal" form:"nominal"`
	Tipe            string `json:"tipe" form:"tipe"`
	Date 			string `json:"date" form:"date"`
	Persetujuan     string `json:"persetujuan" form:"persetujuan"`
	UrlBukti        string `json:"url_bukti" form:"url_bukti"`
	UserID          string `json:"user_id" form:"user_id"`
	User            UserRequest `json:"user,omitempty"`
}

type UserRequest struct {
	ID     string `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Role   string `json:"role" form:"role"`
	Devisi string `json:"devisi" form:"devisi"`
}

type LoginReguest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(login LoginReguest)usernodejs.Login{
	return usernodejs.Login{
		Email: login.Email,
		Password: login.Password,
	}
}
func RequestToEntity(user ReimbursementRequest)reimbusment.ReimbursementEntity{
	return reimbusment.ReimbursementEntity{
		ID:              user.ID,
		Description:     user.Description,
		Status:          user.Status,
		BatasanReimburs: user.BatasanReimburs,
		Nominal:         user.Nominal,
		Tipe:            user.Tipe,
		Date: 			 user.Date,
		Persetujuan:     user.Persetujuan,
		UrlBukti:        user.UrlBukti,
		UserID:          user.UserID,
		User:           UserRequestToEntity(user.User),
	}
}

func UserRequestToEntity(user UserRequest)reimbusment.UserEntity{
	return reimbusment.UserEntity{
		ID:     user.ID,
		Name:   user.Name,
		Role:   user.Role,
		Devisi: user.Devisi,
	}
}