package handler

import "be_golang/klp3/features/target"

type TargetResponse struct {
	ID             string `json:"id,omitempty"`
	KontenTarget   string `json:"konten_target,omitempty"`
	Status         string `json:"status,omitempty"`
	DevisiID       string `json:"devisi_id,omitempty"`
	UserIDPembuat  string `json:"user_id_pembuat,omitempty"`
	UserIDPenerima string `json:"user_id_penerima,omitempty"`
	Due_Date       string `json:"due_date,omitempty"`
	Proofs         string `json:"proofs,omitempty"`
}

func EntityToResponse(entity target.TargetEntity) TargetResponse {
	return TargetResponse{
		ID:             entity.ID,
		KontenTarget:   entity.KontenTarget,
		Status:         entity.Status,
		DevisiID:       entity.DevisiID,
		UserIDPembuat:  entity.UserIDPembuat,
		UserIDPenerima: entity.UserIDPenerima,
		Due_Date:       entity.Due_Date,
		Proofs:         entity.Proofs,
	}
}
