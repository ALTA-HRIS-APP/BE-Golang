package apinodejs

import (
	"time"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct untuk data devisi
type Devisi struct {
	ID        string    `json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Struct untuk data role
type Role struct {
	ID        string    `json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type ResponseDataToken struct {
	Meta MetaInfo `json:"meta"`
	Data DataInfo `json:"data"`
}
type MetaInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type DataInfo struct {
	Token string `json:"token"`
}
type ResponseDataUser struct {
	Meta MetaInfo `json:"meta"`
	Data Pengguna `json:"data"`
}

func MappingToken(token DataInfo) DataInfo {
	return DataInfo{
		Token: token.Token,
	}
}

// Struct untuk data pengguna
type Pengguna struct {
	ID          string    `json:"id"`
	NamaLengkap string    `json:"nama_lengkap"`
	Surel       string    `json:"surel"`
	NoHP        string    `json:"no_hp"`
	Jabatan     string    `json:"jabatan"`
	KataSandi   string    `json:"kata_sandi"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DevisiID    string    `json:"devisiId"`
	RoleID      string    `json:"roleId"`
	Devisi      Devisi    `json:"devisi"`
	Role        Role      `json:"role"`
}

// Struct untuk data utama
type Data struct {
	Meta struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    []Pengguna `json:"data"`
	} `json:"meta"`
}

func ByteToResponseById(pengguna Pengguna) Pengguna {
	return Pengguna{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Surel:       pengguna.Surel,
		NoHP:        pengguna.NoHP,
		Jabatan:     pengguna.Jabatan,
		KataSandi:   pengguna.KataSandi,
		Status:      pengguna.Status,
		CreatedAt:   pengguna.CreatedAt,
		UpdatedAt:   pengguna.UpdatedAt,
		DevisiID:    pengguna.DevisiID,
		RoleID:      pengguna.RoleID,
		Devisi:      Devisipe(pengguna.Devisi),
		Role:        Rolepe(pengguna.Role),
	}
}

func ByteToResponse(pengguna Pengguna) Pengguna {
	return Pengguna{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Surel:       pengguna.Surel,
		NoHP:        pengguna.NoHP,
		Jabatan:     pengguna.Jabatan,
		KataSandi:   pengguna.KataSandi,
		Status:      pengguna.Status,
		CreatedAt:   pengguna.CreatedAt,
		UpdatedAt:   pengguna.UpdatedAt,
		DevisiID:    pengguna.DevisiID,
		RoleID:      pengguna.RoleID,
		Devisi:      Devisipe(pengguna.Devisi),
		Role:        Rolepe(pengguna.Role),
	}
}
func Devisipe(devisi Devisi) Devisi {
	return Devisi{
		ID:        devisi.ID,
		Nama:      devisi.Nama,
		CreatedAt: devisi.CreatedAt,
		UpdatedAt: devisi.UpdatedAt,
	}
}

func Rolepe(devisi Role) Role {
	return Role{
		ID:        devisi.ID,
		Nama:      devisi.Nama,
		CreatedAt: devisi.CreatedAt,
		UpdatedAt: devisi.UpdatedAt,
	}
}

type ExternalDataInterface interface {
	LoginUser(login Login) (string, error)
	GetProfile(token string) (Pengguna, error)
	GetAllUser() ([]Pengguna, error)
	GetUserByID(idUser string) (Pengguna, error)
	// Metode lainnya untuk mengakses API eksternal
}
