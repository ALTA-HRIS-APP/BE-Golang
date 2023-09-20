package handler

import (
	"be_golang/klp3/features/absensi"
	"time"
)

type AbsensiRequest struct {
	ID              string      `json:"id" form:"id"`
	OverTimeMasuk   string      `json:"overtime_masuk" form:"overtime_masuk"`
	OverTimeKeluar  string      `json:"overtime_keluar" form:"overtime_keluar"`
	JamMasuk        string      `json:"jam_masuk" form:"jam_masuk"`
	JamKeluar       string      `json:"jam_keluar" form:"jam_keluar"`
	TanggalSekarang time.Time   `json:"tanggal_sekarang" form:"tanggal_sekarang"`
	CreatedAt       string      `json:"check_in" form:"check_in"`
	UpdateAt        string      `json:"check_out" form:"check_out"`
	UserID          string      `json:"user_id" form:"user_id"`
	User            UserRequest `json:"user,omitempty"`
}

type UserRequest struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"nama_lengkap" form:"nama_lengkap"`
}

func RequestToEntity(user AbsensiRequest) absensi.AbsensiEntity {
	// Parsing tanggal dari string
	createdAt, _ := time.Parse("04:05.000", user.CreatedAt)
	updateAt, _ := time.Parse("04:05.000", user.UpdateAt)

	return absensi.AbsensiEntity{
		ID:              user.ID,
		CreatedAt:       createdAt,
		UpdatedAt:       updateAt,
		DeletedAt:       time.Time{},
		UserID:          user.UserID,
		OverTimeMasuk:   user.OverTimeMasuk,
		OverTimePulang:  user.OverTimeKeluar,
		JamMasuk:        user.JamMasuk,
		JamKeluar:       user.JamKeluar,
		TanggalSekarang: "",
		User:            UserRequestToEntity(user.User),
	}
}
func UserRequestToEntity(user UserRequest) absensi.UserEntity {
	return absensi.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
}
