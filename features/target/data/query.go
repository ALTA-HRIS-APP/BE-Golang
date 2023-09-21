package data

import (
	"be_golang/klp3/features/target"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type targetQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) target.TargetDataInterface {
	return &targetQuery{
		db: database,
	}
}
func (r *targetQuery) GetUserByIDAPI(idUser string) (target.PenggunaEntity, error) {
	// Panggil metode GetUserByID dari externalAPI
	user, err := usernodejs.GetByIdUser(idUser)
	if err != nil {
		log.Printf("Error consume api user: %s", err.Error())
		return target.PenggunaEntity{}, err
	}
	dataUser := UserNodeJsToPengguna(user)
	dataUserEntity := UserPenggunaToEntity(dataUser)
	log.Println("consume api successfully")
	return dataUserEntity, nil
}

// Insert implements target.TargetDataInterface.
func (r *targetQuery) Insert(input target.TargetEntity) (string, error) {
	uuid, err := helper.GenerateUUID()
	if err != nil {
		log.Printf("Error generating UUID: %s", err.Error())
		return "", errors.New("failed genereted uuid")
	}

	newTarget := EntityToModel(input)
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
func (r *targetQuery) SelectAll(param target.QueryParam) (int64, []target.TargetEntity, error) {
	// Initialize variables
	var inputModel []Target
	var totalTarget int64

	// Initial query
	query := r.db

	// Handle searching by description if provided
	if param.SearchKonten != "" {
		query = query.Where("konten_target like ?", "%"+param.SearchKonten+"%")
	}
	if param.SearchStatus != "" {
		query = query.Where("status like ?", "%"+param.SearchStatus+"%")
	}

	// Handle special condition for class dashboard
	if param.ExistOtherPage {
		offset := (param.Page - 1) * param.LimitPerPage
		query = query.Offset(offset).Limit(param.LimitPerPage)
	}

	// Execute the query on the database
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		log.Printf("Error retrieving all targets: %s", tx.Error)
		return 0, nil, errors.New("failed to get all targets")
	}
	totalTarget = tx.RowsAffected
	dataPengguna, err := usernodejs.GetAllUser()
	if err != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	var dataUser []User
	for _, v := range dataPengguna {
		dataUser = append(dataUser, PenggunaToUser(v))
	}
	var userEntity []target.UserEntity
	for _, v := range dataUser {
		userEntity = append(userEntity, UserToEntity(v))
	}
	fmt.Println("user entity", userEntity)
	var targetPengguna []TargetPengguna
	for _, v := range inputModel {
		targetPengguna = append(targetPengguna, ModelToPengguna(v))
	}
	fmt.Println("target", targetPengguna)
	var targetEntity []target.TargetEntity
	for i := 0; i < len(userEntity); i++ {
		for j := 0; j < len(targetPengguna); j++ {
			if userEntity[i].ID == targetPengguna[j].User.ID {
				targetPengguna[j].User = User(userEntity[i])
				targetEntity = append(targetEntity, PenggunaToEntity(targetPengguna[j]))
			}
		}
	}
	// resultTargetSlice := ListModelToEntity(inputModel)
	log.Println("Targets read successfully")
	return totalTarget, targetEntity, nil
}

// SelectAllKaryawan implements target.TargetDataInterface.
func (r *targetQuery) SelectAllKaryawan(idUser string, param target.QueryParam) (int64, []target.TargetEntity, error) {
	var inputModel []Target
	var totalTarget int64

	query := r.db

	// Handle searching by description if provided
	if param.SearchKonten != "" {
		query = query.Where("user_id = ? AND konten_target LIKE ?", idUser, "%"+param.SearchKonten+"%")
	}
	if param.SearchStatus != "" {
		query = query.Where("user_id = ? AND status LIKE ?", idUser, "%"+param.SearchStatus+"%")
	}

	// Special condition for class dashboard
	if param.ExistOtherPage {
		offset := (param.Page - 1) * param.LimitPerPage
		query = query.Where("user_id = ?", idUser).Offset(int(offset)).Limit(param.LimitPerPage)
	}

	// Execute the query on the database
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		log.Printf("Error retrieving all targets: %s", tx.Error)
		return 0, nil, errors.New("failed to get all targets")
	}
	totalTarget = tx.RowsAffected
	dataUser, err := usernodejs.GetByIdUser(idUser)
	if err != nil {
		return 0, nil, err
	}
	pengguna := PenggunaToUser(dataUser)
	userEntity := UserToEntity(pengguna)

	var targetPengguna []TargetPengguna
	for _, v := range inputModel {
		targetPengguna = append(targetPengguna, ModelToPengguna(v))
	}
	var targetEntity []target.TargetEntity
	for _, v := range targetPengguna {
		if v.UserIDPenerima == userEntity.ID {
			v.User = User(userEntity)
			targetEntity = append(targetEntity, PenggunaToEntity(v))
		}

	}
	// resultTargetSlice := ListModelToEntity(inputModel)
	log.Println("Targets read successfully")
	return totalTarget, targetEntity, nil
}

// Select implements target.TargetDataInterface.
func (r *targetQuery) Select(targetID string) (target.TargetEntity, error) {
	var targetData Target

	tx := r.db.Where("id = ?", targetID).First(&targetData)
	if tx.Error != nil {
		log.Printf("Error reading target: %s", tx.Error)
		return target.TargetEntity{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when reading target")
		return target.TargetEntity{}, errors.New("target not found")
	}
	// Mapping target to CoreTarget
	coreTarget := ModelToEntity(targetData)
	log.Println("Target read successfully")
	return coreTarget, nil
}

// Update implements target.TargetDataInterface.
func (r *targetQuery) Update(targetID string, targetData target.TargetEntity) error {
	var target Target
	tx := r.db.Where("id = ?", targetID).First(&target)
	log.Printf("Error reading target by id: %s", tx.Error)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when reading target")
		return errors.New("target not found")
	}

	// Mapping Entity Target to Model
	updatedTarget := EntityToModel(targetData)

	// Perform the update of project data in the database
	tx = r.db.Model(&target).Updates(updatedTarget)
	if tx.Error != nil {
		log.Printf("Error updating target: %s", tx.Error)
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	log.Println("Target updated successfully")
	return nil
}

// Delete implements target.TargetDataInterface.
func (r *targetQuery) Delete(targetID string) error {
	var target Target
	tx := r.db.Where("id = ?", targetID).Delete(&target)
	if tx.Error != nil {
		log.Printf("Error deleting target: %s", tx.Error)
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when deleting target")
		return errors.New("target not found")
	}
	log.Println("Target deleted successfully")
	return nil
}
