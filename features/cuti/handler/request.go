package handler

import (
	"be_golang/klp3/features/cuti"
)

type CutiRequest struct {
	ID           string `json:"id" form:"id"`
	TipeCuti     string `json:"tipe_cuti" form:"tipe_cuti"`
	Status       string `json:"status" form:"status"`
	JumlahCuti   int    `json:"jumlah_cuti" form:"jumlah_cuti"`
	BatasanCuti  int    `json:"batasan_cuti" form:"batasan_cuti"`
	Description  string `json:"description" form:"description"`
	Persetujuan  string `json:"persetujuan" form:"persetujan"`
	StartCuti    string `json:"strat_cuti" form:"start_cuti"`
	EndCuti      string `json:"end_cuti" form:"end_cuti"`
	UrlPendukung string `json:"url_pendukung" form:"url_pendukung"`
	UserID       string `json:"user_id" form:"user_id"`
}

func RequestToEntity(cutii CutiRequest) cuti.CutiEntity {
	return cuti.CutiEntity{
		ID:           cutii.ID,
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
