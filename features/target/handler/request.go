package handler

import (
	"be_golang/klp3/features/target"
)

type TargetRequest struct {
	KontenTarget   string `json:"konten_target,omitempty" form:"konten_target"`
	Status         string `json:"status,omitempty" form:"status"`
	DevisiID       string `json:"devisi_id,omitempty" form:"devisi_id"`
	UserIDPembuat  string `json:"user_id_pembuat,omitempty" form:"user_id_pembuat"`
	UserIDPenerima string `json:"user_id_penerima,omitempty" form:"user_id_penerima"`
	DueDate        string `json:"due_date,omitempty" form:"due_date"`
	Proofs         string `json:"proofs,omitempty" form:"proofs"`
}

func TargetRequestToEntity(req TargetRequest) target.TargetEntity {
	return target.TargetEntity{
		KontenTarget:   req.KontenTarget,
		Status:         req.Status,
		DevisiID:       req.DevisiID,
		UserIDPembuat:  req.UserIDPembuat,
		UserIDPenerima: req.UserIDPenerima,
		DueDate:        req.DueDate,
		Proofs:         req.Proofs,
	}
}
