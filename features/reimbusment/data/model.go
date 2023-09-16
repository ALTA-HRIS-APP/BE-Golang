package data

import (
	"be_golang/klp3/features/reimbusment"
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
	User            User 			`gorm:"foreignKey:UserID"`
}

type User struct{
	ID        		string `gorm:"primaryKey;type:varchar(255)"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Name 			string
	Role 			string
	Devisi 			string	
}

// func PenggunaToUser(user usernodejs.Pengguna)User{
// 	return User{
// 		ID: user.ID,
// 		Name: user.NamaLengkap,
// 		Role: user.Role.Nama,
// 		Devisi: user.Devisi.Nama,
// 	}
// }

func UserEntityToModel(user reimbusment.UserEntity)User{
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Role:      user.Role,
		Devisi:    user.Devisi,
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
		User: UserEntityToModel(user.User),
	}
}
func UserModelToEntity(user User)reimbusment.UserEntity{
	return reimbusment.UserEntity{
		ID:        user.ID,
		Name:      user.Name,
		Role:      user.Role,
		Devisi:    user.Devisi,
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
		User: UserModelToEntity(user.User),
	}
}