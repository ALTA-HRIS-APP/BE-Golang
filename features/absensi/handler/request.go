package handler

import (
	"be_golang/klp3/features/absensi"
)

type AbsensiRequest struct {
	ID             string      `json:"id" form:"id"`
	OverTimeMasuk  string      `json:"overtime_masuk" form:"overtime_masuk"`
	OverTimeKeluar string      `json:"overtime_keluar" form:"overtime_keluar"`
	JamMasuk       string      `json:"jam_masuk" form:"jam_masuk"`
	JamKeluar      string      `json:"jam_keluar" form:"jam_keluar"`
	UserID         string      `json:"user_id" form:"user_id"`
	User           UserRequest `json:"user,omitempty"`
}

type UserRequest struct {
	ID     string `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Role   string `json:"role" form:"role"`
	Devisi string `json:"devisi" form:"devisi"`
}

func RequestToEntity(user AbsensiRequest) absensi.AbsensiEntity {
	return absensi.AbsensiEntity{
		ID:             user.ID,
		UserID:         user.UserID,
		OverTimeMasuk:  user.OverTimeMasuk,
		OverTimePulang: user.OverTimeKeluar,
		JamMasuk:       user.JamMasuk,
		JamKeluar:      user.JamKeluar,
	}
}

func UserRequestToEntity(user UserRequest) absensi.UserEntity {
	return absensi.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
}
