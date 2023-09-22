package data

import (
	"be_golang/klp3/features/target"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"
	"fmt"

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
		return target.PenggunaEntity{}, err
	}
	dataUser := UserNodeJsToPengguna(user)
	dataUserEntity := UserPenggunaToEntity(dataUser)
	return dataUserEntity, nil
}

// Insert implements target.TargetDataInterface.
func (r *targetQuery) Insert(input target.TargetEntity) (string, error) {
	uuid, err := helper.GenerateUUID()
	if err != nil {
		return "", errors.New("failed genereted uuid")
	}

	newTarget := EntityToModel(input)
	newTarget.ID = uuid
	//simpan ke db
	tx := r.db.Create(&newTarget)
	if tx.Error != nil {
		return "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return "", errors.New("target not found")
	}
	return newTarget.ID, nil
}

// SelectAll implements target.TargetDataInterface.
func (r *targetQuery) SelectAll(token string, param target.QueryParam) (int64, []target.TargetEntity, error) {
	// Initialize variables
	var inputModel []Target
	var totalTarget int64

	// Initial query
	query := r.db
	if param.ExistOtherPage {
		offset := (param.Page - 1) * param.LimitPerPage
		fmt.Println("offset", offset)
		if param.SearchKonten != "" {
			query = query.Where("konten_target like ?", "%"+param.SearchKonten+"%")
		}
		if param.SearchStatus != "" {
			query = query.Where("status like ?", "%"+param.SearchStatus+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all target")
		}
		totalTarget = tx.RowsAffected
		query = query.Offset(offset).Limit(param.LimitPerPage)
	}
	// Handle searching by description if provided
	if param.SearchKonten != "" {
		query = query.Where("konten_target like ?", "%"+param.SearchKonten+"%")
	}
	if param.SearchStatus != "" {
		query = query.Where("status like ?", "%"+param.SearchStatus+"%")
	}

	// Execute the query on the database
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("failed to get all targets")
	}
	// dataPengguna, err := usernodejs.GetAllUser(token)
	// if err != nil {
	// 	return 0, nil, err
	// }

	// var dataUser []User
	// for _, v := range dataPengguna {
	// 	dataUser = append(dataUser, PenggunaToUser(v))
	// }

	// var userEntity []target.UserEntity
	// for _, v := range dataUser {
	// 	userEntity = append(userEntity, UserToEntity(v))
	// }

	// var targetPengguna []TargetPengguna
	// for _, v := range inputModel {
	// 	targetPengguna = append(targetPengguna, ModelToPengguna(v))
	// }

	// var targetEntity []target.TargetEntity
	// for i := 0; i < len(userEntity); i++ {
	// 	for j := 0; j < len(targetPengguna); j++ {
	// 		if userEntity[i].ID == targetPengguna[j].UserIDPenerima {
	// 			targetPengguna[j].User = User(userEntity[i])
	// 			targetEntity = append(targetEntity, PenggunaToEntity(targetPengguna[j]))
	// 		}
	// 	}
	// }
	targetEntity := ListModelToEntity(inputModel)
	return totalTarget, targetEntity, nil
}

// SelectAllKaryawan implements target.TargetDataInterface.
func (r *targetQuery) SelectAllKaryawan(idUser string, param target.QueryParam) (int64, []target.TargetEntity, error) {
	var inputModel []Target
	var totalTarget int64

	query := r.db
	if param.ExistOtherPage {
		offset := (param.Page - 1) * param.LimitPerPage
		if param.SearchKonten != "" {
			query = query.Where("user_id=? and konten_target like ?", idUser, "%"+param.SearchKonten+"%")
		}
		if param.SearchStatus != "" {
			query = query.Where("user_id=? and status like ?", idUser, "%"+param.SearchStatus+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all targets")
		}
		totalTarget = tx.RowsAffected
		query = query.Offset(offset).Limit(param.LimitPerPage)
	}
	// Handle searching by description if provided
	if param.SearchKonten != "" {
		query = query.Where("user_id = ? AND konten_target LIKE ?", idUser, "%"+param.SearchKonten+"%")
	}
	if param.SearchStatus != "" {
		query = query.Where("user_id = ? AND status LIKE ?", idUser, "%"+param.SearchStatus+"%")
	}

	// Execute the query on the database
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("failed to get all targets")
	}
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

	fmt.Printf("Number of targets retrieved for user %s: %d\n", idUser, len(targetPengguna))

	var targetEntity []target.TargetEntity
	for _, v := range targetPengguna {
		if v.UserIDPenerima == userEntity.ID {
			v.User = User(userEntity)
			targetEntity = append(targetEntity, PenggunaToEntity(v))

			fmt.Printf("Target matched: UserIDPenerima: %s, userEntity.ID: %s\n", v.UserIDPenerima, userEntity.ID)
		}
	}

	// resultTargetSlice := ListModelToEntity(inputModel)
	return totalTarget, targetEntity, nil
}

// Select implements target.TargetDataInterface.
func (r *targetQuery) Select(idTarget string) (target.TargetEntity, error) {
	var targetData Target

	tx := r.db.Where("id = ?", idTarget).First(&targetData)
	if tx.Error != nil {
		return target.TargetEntity{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return target.TargetEntity{}, errors.New("target not found")
	}
	// Mapping target to CoreTarget
	coreTarget := ModelToEntity(targetData)
	return coreTarget, nil
}

// Update implements target.TargetDataInterface.
func (r *targetQuery) Update(idTarget string, targetData target.TargetEntity) error {
	var target Target
	tx := r.db.Where("id = ?", idTarget).First(&target)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	// Mapping Entity Target to Model
	updatedTarget := EntityToModel(targetData)

	// Perform the update of project data in the database
	tx = r.db.Model(&target).Updates(updatedTarget)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}

// Delete implements target.TargetDataInterface.
func (r *targetQuery) Delete(idTarget string) error {
	var target Target
	tx := r.db.Where("id = ?", idTarget).Delete(&target)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}
	return nil
}
