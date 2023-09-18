package data

import (
	"be_golang/klp3/features/target"
	"time"

	"gorm.io/gorm"
)

type Target struct {
	ID             string `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	KontenTarget   string         `gorm:"column:konten_target;not null"`
	Status         string         `gorm:"type:enum('not completed','completed');column:status;"`
	DevisiID       string         `gorm:"column:devisi_id;not null"`
	UserIDPembuat  string         `gorm:"column:user_id_pembuat;not null"`
	UserIDPenerima string         `gorm:"column:user_id_penerima;not null"`
	Due_Date       string         `gorm:"column:due_date;not null"`
	Proofs         string         `gorm:"column:proofs"`
}

func MapEntityToModel(entity target.TargetEntity) Target {
	return Target{
		ID:             entity.ID,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
		KontenTarget:   entity.KontenTarget,
		Status:         entity.Status,
		DevisiID:       entity.DevisiID,
		UserIDPembuat:  entity.UserIDPembuat,
		UserIDPenerima: entity.UserIDPenerima,
		Due_Date:       entity.Due_Date,
		Proofs:         entity.Proofs,
	}
}

func MapModelToEntity(model Target) target.TargetEntity {
	return target.TargetEntity{
		ID:             model.ID,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
		DeletedAt:      model.DeletedAt.Time,
		KontenTarget:   model.KontenTarget,
		Status:         model.Status,
		DevisiID:       model.DevisiID,
		UserIDPembuat:  model.UserIDPembuat,
		UserIDPenerima: model.UserIDPenerima,
		Due_Date:       model.Due_Date,
		Proofs:         model.Proofs,
	}
}

func ListModelToEntity(models []Target) []target.TargetEntity {
	var TargetsEntity []target.TargetEntity
	for _, model := range models {
		TargetsEntity = append(TargetsEntity, target.TargetEntity{
			ID:             model.ID,
			CreatedAt:      model.CreatedAt,
			UpdatedAt:      model.UpdatedAt,
			DeletedAt:      model.DeletedAt.Time,
			KontenTarget:   model.KontenTarget,
			Status:         model.Status,
			DevisiID:       model.DevisiID,
			UserIDPembuat:  model.UserIDPembuat,
			UserIDPenerima: model.UserIDPenerima,
			Due_Date:       model.Due_Date,
			Proofs:         model.Proofs,
		})
	}
	return TargetsEntity
}
