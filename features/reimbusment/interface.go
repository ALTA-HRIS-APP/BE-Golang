package reimbusment

import (
	"time"
)

type ReimbursementEntity struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	Description     string 	`validate:"required"`
	Status          string 	
	BatasanReimburs int 
	Nominal         int 	`validate:"required"`
	Tipe            string 	`validate:"required"`
	Date 			string `validate:"required"`
	Persetujuan     string 
	UrlBukti        string 	`validate:"required"`
	UserID          string 
	User 			UserEntity
}

type UserEntity struct{
	ID        		string 
	Name 			string
}
type PenggunaEntity struct{
	ID          string   
	NamaLengkap string    
	Jabatan     string   
	Devisi      string 	
}

type QueryParams struct {
	Page            int
	ItemsPerPage    int
	SearchName      string
	IsClassDashboard bool
}

type ReimbusmentDataInterface interface{
	Insert(input ReimbursementEntity)(error)
	UpdateKaryawan(input ReimbursementEntity,id string)(error)
	Update(input ReimbursementEntity,id string)(error)
	SelectById(id string)(ReimbursementEntity,error)
	SelectAllKaryawan(idUser string,param QueryParams)(int64,[]ReimbursementEntity,error)
	SelectAll(token string,param QueryParams)(int64,[]ReimbursementEntity,error)
	Delete(id string)error
	SelectUserById(idUser string)(PenggunaEntity,error)
}

type ReimbusmentServiceInterface interface{
	Add(input ReimbursementEntity)(error)
	Edit(input ReimbursementEntity,id string,idUser string)(error)
	Get(token string,idUser string,param QueryParams)(bool,[]ReimbursementEntity,error)
	Delete(id string)error
	GetReimbusherById(id string)(ReimbursementEntity,error)
}