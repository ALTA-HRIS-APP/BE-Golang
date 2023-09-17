package handler

import "be_golang/klp3/features/reimbusment"

type ReimbursementResponse struct {
	ID              string `json:"id,omitempty"`
	Description     string `json:"description,omitempty"`
	Status          string `json:"status,omitempty"`
	BatasanReimburs int `json:"batasan_reimburs,omitempty"`
	Nominal         int `json:"nominal,omitempty"`
	Tipe            string `json:"tipe,omitempty"`
	Date 			string `json:"date,omitempty"`
	Persetujuan     string `json:"persetujuan,omitempty"`
	UrlBukti        string `json:"url_bukti,omitempty"`
	UserID          string `json:"user_id,omitempty"`
	User            UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func EntityToResponse(user reimbusment.ReimbursementEntity) ReimbursementResponse {
	return ReimbursementResponse{
		ID:              user.ID,
		Description:     user.Description,
		Status:          user.Status,
		BatasanReimburs: user.BatasanReimburs,
		Nominal:         user.Nominal,
		Tipe:            user.Tipe,
		Persetujuan:     user.Persetujuan,
		UrlBukti:        user.UrlBukti,
		UserID:          user.UserID,
		User:            UserEntityToResponse(user.User),
	}
}

func UserEntityToResponse(user reimbusment.UserEntity) UserResponse {
	return UserResponse{
		ID:     user.ID,
		Name:   user.Name,
	}
}