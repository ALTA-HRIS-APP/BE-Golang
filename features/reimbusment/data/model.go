package data

import (
	"be_golang/klp3/features/reimbusment"
	usernodejs "be_golang/klp3/features/userNodejs"
	"time"

	"gorm.io/gorm"
)

type Reimbursement struct {
	ID        		string 			`gorm:"primaryKey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Description 	string
	Status 			string 			`gorm:"default:pending"`
	BatasanReimburs int 			`gorm:"default:5000000"`
	Nominal 		int
	Date 			string           `gorm:"column:date;not nul"`
	Tipe 			string
	Persetujuan 	string 			`gorm:"default:-"`
	UrlBukti 		string
	UserID 			string 			`gorm:"type:varchar(255)"`
}

type ReimbursementPengguna struct {
	ID        		string 			
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 
	Description 	string
	Status 			string 			
	BatasanReimburs int 			
	Nominal 		int
	Date 			string           
	Tipe 			string
	Persetujuan 	string 			
	UrlBukti 		string
	UserID 			string
	User            User 			
}

type User struct{
	ID        		string 
	Name 			string	
}

type Pengguna struct{
	ID          string    
	NamaLengkap string    
	Jabatan     string 
	Devisi 		string   	
}

func UserNodeJskePengguna(pengguna usernodejs.Pengguna)Pengguna{
	return Pengguna{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Jabatan:     pengguna.Jabatan,
		Devisi: 	 pengguna.Devisi.Nama,
	}
}
func UserPenggunaToEntity(pengguna Pengguna)reimbusment.PenggunaEntity{
	return reimbusment.PenggunaEntity{
		ID:          pengguna.ID,
		NamaLengkap: pengguna.NamaLengkap,
		Jabatan:     pengguna.Jabatan,
		Devisi: 	 pengguna.Devisi,
	}
}

func ModelToPengguna(user Reimbursement)ReimbursementPengguna{
	return ReimbursementPengguna{
		ID:              user.ID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
		Description:     user.Description,
		Status:          user.Status,
		BatasanReimburs: user.BatasanReimburs,
		Nominal:         user.Nominal,
		Date:            user.Date,
		Tipe:            user.Tipe,
		Persetujuan:     user.Persetujuan,
		UrlBukti:        user.UrlBukti,
		UserID:          user.UserID,
	}
}

func PenggunaToEntity(user ReimbursementPengguna)reimbusment.ReimbursementEntity{
	return reimbusment.ReimbursementEntity{
		ID:              user.ID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		Description:     user.Description,
		Status:          user.Status,
		BatasanReimburs: user.BatasanReimburs,
		Nominal:         user.Nominal,
		Date:            user.Date,
		Tipe:            user.Tipe,
		Persetujuan:     user.Persetujuan,
		UrlBukti:        user.UrlBukti,
		UserID:          user.UserID,
		User: UserToEntity(user.User),
	}
}

func PenggunaToUser(user usernodejs.Pengguna)User{
	return User{
		ID: user.ID,
		Name: user.NamaLengkap,
	}
}

func UserToEntity(user User)reimbusment.UserEntity{
	return reimbusment.UserEntity{
		ID:   user.ID,
		Name: user.Name,
	}
}

func EntityToModel(user reimbusment.ReimbursementEntity)Reimbursement{
	return Reimbursement{
		ID:              user.ID,
		Description:     user.Description,
		Status:          user.Status,
		BatasanReimburs: user.BatasanReimburs,
		Nominal:         user.Nominal,
		Tipe:            user.Tipe,
		Date: 			 user.Date,
		Persetujuan:     user.Persetujuan,
		UrlBukti:        user.UrlBukti,
		UserID:          user.UserID,
	}
}


func ModelToEntity(user Reimbursement)reimbusment.ReimbursementEntity{
	return reimbusment.ReimbursementEntity{
		ID:              user.ID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt.Time,
		Description:     user.Description,
		Status:          user.Status,
		BatasanReimburs: user.BatasanReimburs,
		Nominal:         user.Nominal,
		Tipe:            user.Tipe,
		Date: 			 user.Date,
		Persetujuan:     user.Persetujuan,
		UrlBukti:        user.UrlBukti,
		UserID:          user.UserID,
	}
}