package handler

import (
	"be_golang/klp3/features/cuti"
	"time"
)

type CutiResponse struct {
	ID           string    `json:"id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	TipeCuti     string    `json:"tipe_cuti,omitempty"`
	Status       string    `json:"status,omitempty"`
	JumlahCuti   int       `json:"jumlah_cuti,omitempty"`
	BatasanCuti  int       `json:"batasan_cuti,omitempty"`
	Description  string    `json:"description,omitempty"`
	Persetujuan  string    `json:"persetujuan,omitempty"`
	StartCuti    string    `json:"strat_cuti,omitempty"`
	EndCuti      string    `json:"end_cuti,omitempty"`
	UrlPendukung string    `json:"url_pendukung,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
}

func EntityToResponse(cutii cuti.CutiEntity) CutiResponse {
	return CutiResponse{
		ID:           cutii.ID,
		CreatedAt:    cutii.CreatedAt,
		UpdatedAt:    cutii.UpdatedAt,
		TipeCuti:     cutii.TipeCuti,
		Status:       cutii.Status,
		JumlahCuti:   cutii.JumlahCuti,
		BatasanCuti:  cutii.BatasanCuti,
		Description:  cutii.Description,
		Persetujuan:  cutii.Persetujuan,
		StartCuti:    cutii.StartCuti,
		EndCuti:      cutii.EndCuti,
		UrlPendukung: cutii.UrlPendukung,
		UserID:       cutii.UserID,
	}
}
