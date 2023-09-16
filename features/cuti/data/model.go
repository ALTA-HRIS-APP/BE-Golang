package data

import (
	"be_golang/klp3/features/cuti"
	"time"

	"gorm.io/gorm"
)

type Cuti struct {
	ID           string `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	TipeCuti     string         `gorm:"type:enum('melahirkan','hari raya','tahunan');default:'tahunan';column:tipe_cuti"`
	Status       string
	JumlahCuti   int
	BatasanCuti  int `gorm:"column:batasan_cuti;default:90"`
	Description  string
	Persetujuan  string
	StartCuti    string
	EndCuti      string
	UrlPendukung string
	UserID       string
}

func EntityToModel(cutii cuti.CutiEntity) Cuti {
	return Cuti{
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
func ModelToEntity(cutii Cuti) cuti.CutiEntity {
	return cuti.CutiEntity{
		ID:           cutii.ID,
		CreatedAt:    cutii.CreatedAt,
		UpdatedAt:    cutii.UpdatedAt,
		DeletedAt:    cutii.DeletedAt.Time,
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
