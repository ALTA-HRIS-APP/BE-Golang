package data

import (
	"be_golang/klp3/features/cuti"
	usernodejs "be_golang/klp3/features/userNodejs"
	"time"

	"gorm.io/gorm"
)

type Cuti struct {
	ID           string `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	TipeCuti     string         `gorm:"type:enum('melahirkan','hari raya','tahunan');default:'tahunan';column:tipe_cuti"`
	Status       string         `gorm:"default:pending"`
	JumlahCuti   int
	BatasanCuti  int `gorm:"column:batasan_cuti;default:90"`
	Description  string
	Persetujuan  string `gorm:"default:pending"`
	StartCuti    string
	EndCuti      string
	UrlPendukung string
	UserID       string
}

type CutiPengguna struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	TipeCuti     string
	Status       string
	JumlahCuti   int
	BatasanCuti  int
	Description  string
	Persetujuan  string
	StartCuti    string
	EndCuti      string
	UrlPendukung string
	UserID       string
	User         User
}

type User struct {
	ID   string
	Name string
}

func ModelToPengguna(user Cuti) CutiPengguna {
	return CutiPengguna{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
		TipeCuti:     user.TipeCuti,
		Status:       user.Status,
		JumlahCuti:   user.JumlahCuti,
		BatasanCuti:  user.BatasanCuti,
		Description:  user.Description,
		Persetujuan:  user.Persetujuan,
		StartCuti:    user.StartCuti,
		EndCuti:      user.EndCuti,
		UrlPendukung: user.UrlPendukung,
		UserID:       user.UserID,
	}
}

func PengunaToEntity(user CutiPengguna) cuti.CutiEntity {
	return cuti.CutiEntity{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		TipeCuti:     user.TipeCuti,
		Status:       user.Status,
		JumlahCuti:   user.JumlahCuti,
		BatasanCuti:  user.BatasanCuti,
		Description:  user.Description,
		Persetujuan:  user.Persetujuan,
		StartCuti:    user.StartCuti,
		EndCuti:      user.EndCuti,
		UrlPendukung: user.UrlPendukung,
		UserID:       user.UserID,
		User:         UserToEntity(user.User),
	}
}

func PenggunaToUser(user usernodejs.Pengguna) User {
	return User{
		ID:   user.ID,
		Name: user.NamaLengkap,
	}
}

func UserToEntity(user User) cuti.UserEntity {
	return cuti.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
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
