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
	Persetujuan     string 
	UrlBukti        string 	`validate:"required"`
	UserID          string 
	User 			UserEntity
}

type UserEntity struct{
	ID        		string 
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		time.Time
	Name 			string
	Role 			string
	Devisi 			string	
}

type ReimbusmentDataInterface interface{
	Insert(input ReimbursementEntity)(string,error)
	SelectUser(UserID string)(UserEntity,error)
}

type ReimbusmentServiceInterface interface{
	Add(input ReimbursementEntity)(error)
}