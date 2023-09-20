package handler

import (
	"be_golang/klp3/features/absensi"
	"time"
)

type AbsensiResponse struct {
	ID              string       `json:"id,omitempty"`
	OverTimeMasuk   string       `json:"overtime_masuk" form:"overtime_masuk"`
	JamMasuk        string       `json:"jam_masuk" form:"jam_masuk"`
	JamKeluar       string       `json:"jam_keluar" form:"jam_keluar"`
	TanggalSekarang string       `json:"tanggal_sekarang" form:"tanggal_sekarang"`
	CreatedAt       time.Time    `json:"check_in" form:"check_in"`
	UpdateAt        time.Time    `json:"check_out" form:"check_out"`
	UserID          string       `json:"user_id,omitempty"`
	User            UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func EntityToResponse(user absensi.AbsensiEntity) AbsensiResponse {
	return AbsensiResponse{
		ID:              user.ID,
		OverTimeMasuk:   user.OverTimeMasuk,
		JamMasuk:        user.JamMasuk,
		JamKeluar:       user.JamKeluar,
		TanggalSekarang: user.TanggalSekarang,
		CreatedAt:       user.CreatedAt,
		UpdateAt:        user.UpdatedAt,
		UserID:          user.UserID,
		User:            UserEntityToResponse(user.User),
	}
}

func UserEntityToResponse(user absensi.UserEntity) UserResponse {
	return UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}
