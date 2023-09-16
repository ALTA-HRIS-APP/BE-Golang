package data

import (
	"be_golang/klp3/features/target"
	"be_golang/klp3/helper"
	"errors"
	"log"

	"gorm.io/gorm"
)

type targetQuery struct {
	db *gorm.DB
}

// Insert implements target.TargetDataInterface.
func (r *targetQuery) Insert(input target.TargetEntity) (string, error) {
	uuid, err := helper.GenerateUUID()
	if err != nil {
		log.Printf("Error generating UUID: %s", err.Error())
		return "", errors.New("failed genereted uuid")
	}
	newTarget := MapEntityToModel(input)
	newTarget.ID = uuid
	//simpan ke db
	tx := r.db.Create(&newTarget)
	if tx.Error != nil {
		log.Printf("Error inserting target: %s", tx.Error)
		return "", tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when inserting target")
		return "", errors.New("target not found")
	}
	log.Println("Target inserted successfully")
	return newTarget.ID, nil
}

func New(database *gorm.DB) target.TargetDataInterface {
	return &targetQuery{
		db: database,
	}
}
