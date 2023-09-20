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
	CreatedAt       string       `json:"check_in" form:"check_in"`
	UpdateAt        string       `json:"check_out" form:"check_out"`
	UserID          string       `json:"user_id,omitempty"`
	User            UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"nama_lengkap,omitempty"`
}

func EntityToResponse(user absensi.AbsensiEntity) AbsensiResponse {
	// Parsing tanggal dan jam dari string
	createdAt, _ := time.Parse("2006-01-02T15:04:05.000Z", user.CreatedAt.Format("15:04:05.000"))
	updateAt, _ := time.Parse("2006-01-02T15:04:05.000Z", user.UpdatedAt.Format("15:04:05.000"))

	// Membuat string terpisah untuk tanggal dan jam

	createdTimeStr := createdAt.Format("15:04:05.000")
	TanggalSekarangg := updateAt.Format("2006-01-02")
	updateTimeStr := updateAt.Format("15:04:05.000")

	return AbsensiResponse{
		ID:              user.ID,
		OverTimeMasuk:   user.OverTimeMasuk,
		JamMasuk:        user.JamMasuk,
		JamKeluar:       user.JamKeluar,
		TanggalSekarang: TanggalSekarangg,
		CreatedAt:       createdTimeStr,
		UpdateAt:        updateTimeStr,
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
