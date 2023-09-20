package data

import (
	"be_golang/klp3/features/externalapi"
	"be_golang/klp3/features/target"
	"be_golang/klp3/helper"
	"errors"
	"log"

	"gorm.io/gorm"
)

type targetQuery struct {
	db          *gorm.DB
	externalAPI externalapi.ExternalDataInterface
}

func New(database *gorm.DB, externalAPI externalapi.ExternalDataInterface) target.TargetDataInterface {
	return &targetQuery{
		db:          database,
		externalAPI: externalAPI,
	}
}

func (r *targetQuery) GetUserByIDFromExternalAPI(idUser string) (externalapi.Pengguna, error) {
	// Panggil metode GetUserByID dari externalAPI
	user, err := r.externalAPI.GetUserByID(idUser)
	if err != nil {
		log.Printf("Error consume api user: %s", err.Error())
		return externalapi.Pengguna{}, err
	}
	log.Println("consume api successfully")
	return user, nil
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

// SelectAll implements target.TargetDataInterface.
func (r *targetQuery) SelectAll(userID string, param target.QueryParam) (int64, []target.TargetEntity, error) {
	// Inisialisasi variabel
	var inputModel []Target
	var totalTarget int64

	// Query awal
	query := r.db.Where("user_id = ?", userID)

	// Handle pencarian berdasarkan nama jika diberikan
	if param.SearchKonten != "" {
		query = query.Where("description like ?", "%"+param.SearchKonten+"%")
	}
	if param.SearchStatus != "" {
		query = query.Where("description like ?", "%"+param.SearchStatus+"%")
	}

	// Handle kondisi khusus untuk dashboard kelas
	if param.ExistOtherPage {
		offset := (param.Page - 1) * param.LimitPerPage
		query = query.Offset(offset).Limit(param.LimitPerPage)
	}

	// Eksekusi query ke database
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("failed to get all target")
	}
	totalTarget = tx.RowsAffected

	resultTargetSlice := ListModelToEntity(inputModel)
	log.Println("Read target successfully")
	return totalTarget, resultTargetSlice, nil

}

// Select implements target.TargetDataInterface.
func (r *targetQuery) Select(targetID string, userID string) (target.TargetEntity, error) {
	var targetData Target

	tx := r.db.Where("id = ? AND user_id = ?", targetID, userID).First(&targetData)
	if tx.Error != nil {
		log.Printf("Error read target: %s", tx.Error)
		return target.TargetEntity{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when read target")
		return target.TargetEntity{}, errors.New("target not found")
	}
	//Mapping target to CorePtarget
	coreTarget := MapModelToEntity(targetData)
	log.Println("Read target successfully")
	return coreTarget, nil
}

// Update implements target.TargetDataInterface.
func (r *targetQuery) Update(targetID string, userID string, targetData target.TargetEntity) error {
	var target Target
	tx := r.db.Where("id = ? AND user_id = ?", targetID, userID).First(&target)
	log.Printf("Error read id: %s", tx.Error)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when read target")
		return errors.New("target not found")
	}

	//Mapping Entity Target to Model
	updatedTarget := MapEntityToModel(targetData)

	// Lakukan pembaruan data proyek dalam database
	tx = r.db.Model(&target).Updates(updatedTarget)
	if tx.Error != nil {
		log.Printf("Error update target: %s", tx.Error)
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	log.Println("Update target successfully")
	return nil
}

// Delete implements target.TargetDataInterface.
func (r *targetQuery) Delete(targetID string, userID string) error {
	var target Target
	tx := r.db.Where("id = ? AND user_id = ?", targetID, userID).Delete(&target)
	if tx.Error != nil {
		log.Printf("Error read id: %s", tx.Error)
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when read target")
		return errors.New("target not found")
	}
	log.Println("Update target successfully")
	return nil
}
