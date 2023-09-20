package data

import (
	"be_golang/klp3/features/absensi"
	usernodejs "be_golang/klp3/features/userNodejs"
	"time"

	"gorm.io/gorm"
)

type Absensi struct {
	ID             string `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserID         string
	OverTimeMasuk  string
	OverTimePulang string
	JamMasuk       string
	JamKeluar      string
}

type User struct {
	ID   string
	Name string
}

type Pengguna struct {
	ID          string
	NamaLengkap string
	Jabatan     string
}

type AbsensiPengguna struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	UserID          string
	OverTimeMasuk   string
	OverTimePulang  string
	JamMasuk        string
	JamKeluar       string
	User            User
	TanggalSekarang time.Time
}

func UserNodeJskePengguna(pengguna usernodejs.Pengguna) Pengguna {
	return Pengguna{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Jabatan:     pengguna.Jabatan,
	}
}

func EntityToModel(absen absensi.AbsensiEntity) Absensi {
	return Absensi{
		ID:             absen.ID,
		UserID:         absen.UserID,
		OverTimeMasuk:  absen.OverTimeMasuk,
		OverTimePulang: absen.OverTimePulang,
		JamMasuk:       absen.JamMasuk,
		JamKeluar:      absen.JamKeluar,
	}
}

func ModelToEntity(user Absensi) absensi.AbsensiEntity {
	return absensi.AbsensiEntity{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		UserID:         user.UserID,
		OverTimeMasuk:  user.OverTimeMasuk,
		OverTimePulang: user.OverTimePulang,
		JamMasuk:       user.JamMasuk,
		JamKeluar:      user.JamKeluar,
		User:           absensi.UserEntity{},
	}
}

func UserPenggunaToEntity(pengguna Pengguna) absensi.PenggunaEntity {
	return absensi.PenggunaEntity{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Jabatan:     pengguna.Jabatan,
	}
}

func UserModelToEntity(user User) absensi.UserEntity {
	return absensi.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
}
func PenggunaToEntity(user AbsensiPengguna) absensi.AbsensiEntity {
	return absensi.AbsensiEntity{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt.Time,
		UserID:         user.UserID,
		OverTimeMasuk:  user.OverTimeMasuk,
		OverTimePulang: user.OverTimePulang,
		JamMasuk:       user.JamMasuk,
		JamKeluar:      user.JamKeluar,
	}
}

func PenggunaToUser(user usernodejs.Pengguna) User {
	return User{
		ID:   user.ID,
		Name: user.NamaLengkap,
	}
}

func UserToEntity(user User) absensi.UserEntity {
	return absensi.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
}

func ModelToPengguna(user Absensi) AbsensiPengguna {
	return AbsensiPengguna{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
		UserID:         user.UserID,
		OverTimeMasuk:  user.OverTimeMasuk,
		OverTimePulang: user.OverTimePulang,
		JamMasuk:       user.JamMasuk,
		JamKeluar:      user.JamKeluar,
	}
}
