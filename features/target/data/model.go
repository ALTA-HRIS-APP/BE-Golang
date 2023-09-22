package data

import (
	"be_golang/klp3/features/target"
	usernodejs "be_golang/klp3/features/userNodejs"
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
	DueDate        string         `gorm:"column:due_date;not null"`
	Proofs         string         `gorm:"column:proofs"`
}

type TargetPengguna struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	KontenTarget   string
	Status         string
	DevisiID       string
	UserIDPembuat  string
	UserIDPenerima string
	DueDate        string
	Proofs         string
	User           User
}

type User struct {
	ID   string
	Name string
}

type Pengguna struct {
	ID          string
	NamaLengkap string
	Jabatan     string
	Devisi      string
}

func UserNodeJsToPengguna(nodejs usernodejs.Pengguna) Pengguna {
	return Pengguna{
		ID:          nodejs.ID,
		NamaLengkap: nodejs.NamaLengkap,
		Jabatan:     nodejs.Jabatan,
		Devisi:      nodejs.Devisi.Nama,
	}
}

func UserPenggunaToEntity(pengguna Pengguna) target.PenggunaEntity {
	return target.PenggunaEntity{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Jabatan:     pengguna.Jabatan,
		Devisi:      pengguna.Devisi,
	}
}

// Model target ke Target pengguna
func ModelToPengguna(user Target) TargetPengguna {
	return TargetPengguna{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		KontenTarget:   user.KontenTarget,
		Status:         user.Status,
		DevisiID:       user.DevisiID,
		UserIDPembuat:  user.UserIDPembuat,
		UserIDPenerima: user.UserIDPenerima,
		DueDate:        user.DueDate,
		Proofs:         user.Proofs,
	}
}

// Target Pengguna ke entity target
func PenggunaToEntity(user TargetPengguna) target.TargetEntity {
	return target.TargetEntity{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt.Time,
		KontenTarget:   user.KontenTarget,
		Status:         user.Status,
		DevisiID:       user.DevisiID,
		UserIDPembuat:  user.UserIDPembuat,
		UserIDPenerima: user.UserIDPenerima,
		DueDate:        user.DueDate,
		Proofs:         user.Proofs,
		User:           UserToEntity(user.User),
	}
}

func PenggunaToUser(user usernodejs.Pengguna) User {
	return User{
		ID:   user.ID,
		Name: user.NamaLengkap,
	}
}

func UserToEntity(user User) target.UserEntity {
	return target.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
}

func EntityToModel(target target.TargetEntity) Target {
	return Target{
		ID:             target.ID,
		CreatedAt:      target.CreatedAt,
		UpdatedAt:      target.UpdatedAt,
		KontenTarget:   target.KontenTarget,
		Status:         target.Status,
		DevisiID:       target.DevisiID,
		UserIDPembuat:  target.UserIDPembuat,
		UserIDPenerima: target.UserIDPenerima,
		DueDate:        target.DueDate,
		Proofs:         target.Proofs,
	}
}
func ModelToEntity(model Target) target.TargetEntity {
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
		DueDate:        model.DueDate,
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
			DueDate:        model.DueDate,
			Proofs:         model.Proofs,
		})
	}
	return TargetsEntity
}
