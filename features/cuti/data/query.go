package data

import (
	"be_golang/klp3/features/cuti"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CutiData struct {
	db *gorm.DB
}

// SelectAllKaryawan implements cuti.CutiDataInterface.
func (repo *CutiData) SelectAllKaryawan(idUser string, param cuti.QueryParams) (int64, []cuti.CutiEntity, error) {
	var inputModel []Cuti
	var total_cuti int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		if param.SearchName != "" {
			query = query.Where("user_id=? and description like ?", idUser, "%"+param.SearchName+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all cuti")
		}
		total_cuti = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	if param.SearchName != "" {
		query = query.Where("user_id=? and description like ?", idUser, "%"+param.SearchName+"%")
	}
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all cuti karyawan")
	}

	dataUser, errUser := usernodejs.GetByIdUser(idUser)
	if errUser != nil {
		return 0, nil, errUser
	}
	pengguna := PenggunaToUser(dataUser)
	userEntity := UserToEntity(pengguna)

	var cutiPengguna []CutiPengguna
	for _, value := range inputModel {
		cutiPengguna = append(cutiPengguna, ModelToPengguna(value))
	}
	var cutiEntity []cuti.CutiEntity
	for _, value := range cutiPengguna {
		if value.UserID == userEntity.ID {
			value.User = User(userEntity)
			cutiEntity = append(cutiEntity, PengunaToEntity(value))
		}
	}
	return total_cuti, cutiEntity, nil
}

// SelectUserById implements cuti.CutiDataInterface.
func (repo *CutiData) SelectUserById(idUser string) (cuti.PenggunaEntity, error) {
	data, err := usernodejs.GetByIdUser(idUser)
	if err != nil {
		return cuti.PenggunaEntity{}, errors.New("error select user by id")
	}
	dataUser := UserNodeJskePengguna(data)
	dataEntity := UserPenggunaToEntity(dataUser)
	return dataEntity, nil
}

// Delete implements cuti.CutiDataInterface.
func (repo *CutiData) Delete(id string) error {
	var inputModel Cuti
	tx := repo.db.Where("id=?", id).Delete(&inputModel)
	if tx.Error != nil {
		return errors.New("delete error cuti")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// SelectById implements cuti.CutiDataInterface.
func (repo *CutiData) SelectById(id string) (cuti.CutiEntity, error) {
	var inputModel Cuti
	tx := repo.db.Where("id=?", id).First(&inputModel)
	if tx.Error != nil {
		return cuti.CutiEntity{}, errors.New("failed get cuti by id")
	}
	output := ModelToEntity(inputModel)
	return output, nil
}

// Update implements cuti.CutiDataInterface.
func (repo *CutiData) Update(input cuti.CutiEntity, id string) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Cuti{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("failed update cuti by id")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// UpdateKaryawan implements cuti.CutiDataInterface.
func (repo *CutiData) UpdateKaryawan(input cuti.CutiEntity, id string) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Cuti{}).Where("id=? and user_id=?", id, input.UserID).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("failed update cuti by id")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// SelectAll implements cuti.CutiDataInterface.
func (repo *CutiData) SelectAll(param cuti.QueryParams) (int64, []cuti.CutiEntity, error) {
	var inputModel []Cuti
	var total_cuti int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		fmt.Println("offset", offset)
		if param.SearchName != "" {
			query = query.Where("description like ?", "%"+param.SearchName+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all cuti")
		}
		total_cuti = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	if param.SearchName != "" {
		query = query.Where("description like ?", "%"+param.SearchName+"%")
	}
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all cuti")
	}

	dataPengguna, errUser := usernodejs.GetAllUser()
	if errUser != nil {
		return 0, nil, errUser
	}
	var dataUser []User
	for _, value := range dataPengguna {
		dataUser = append(dataUser, PenggunaToUser(value))
	}
	var userEntity []cuti.UserEntity
	for _, value := range dataUser {
		userEntity = append(userEntity, UserToEntity(value))
	}
	fmt.Println("user entity", userEntity)
	var reimbushPengguna []CutiPengguna
	for _, value := range inputModel {
		reimbushPengguna = append(reimbushPengguna, ModelToPengguna(value))
	}
	fmt.Println("reimb", reimbushPengguna)
	var reimbushEntity []cuti.CutiEntity
	for i := 0; i < len(userEntity); i++ {
		for j := 0; j < len(reimbushPengguna); j++ {
			if userEntity[i].ID == reimbushPengguna[j].UserID {
				reimbushPengguna[j].User = User(userEntity[i])
				reimbushEntity = append(reimbushEntity, PengunaToEntity(reimbushPengguna[j]))
			}
		}
	}
	return total_cuti, reimbushEntity, nil
}

// SelectAllKaryawan implements cuti.CutiDataInterface.

// insert implements cuti.CutiDataInterface.
func (repo *CutiData) Insert(input cuti.CutiEntity) error {
	inputModel := EntityToModel(input)
	id, err := helper.GenerateUUID()
	if err != nil {
		return errors.New("error create uuid")
	}
	inputModel.ID = id
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("create error")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affective")
	}
	return nil
}

func New(db *gorm.DB) cuti.CutiDataInterface {
	return &CutiData{
		db: db,
	}
}
