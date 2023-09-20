package data

import (
	"be_golang/klp3/features/cuti"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"

	"gorm.io/gorm"
)

type CutiData struct {
	db *gorm.DB
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
func (repo *CutiData) SelectAll() ([]cuti.CutiEntity, error) {
	var inputModel []Cuti
	//offset := (page - 1) * item
	tx := repo.db.Find(&inputModel)
	if tx.Error != nil {
		return nil, errors.New("error get all cuti")
	}

	dataPengguna, errUser := usernodejs.GetAllUser()
	if errUser != nil {
		return nil, errUser
	}
	var dataUser []User
	for _, value := range dataPengguna {
		dataUser = append(dataUser, PenggunaToUser(value))
	}
	var userEntity []cuti.UserEntity
	for _, value := range dataUser {
		userEntity = append(userEntity, UserToEntity(value))
	}
	var cutiPengguna []CutiPengguna
	for _, value := range inputModel {
		cutiPengguna = append(cutiPengguna, ModelToPengguna(value))
	}

	var cutiEntity []cuti.CutiEntity
	for i := 0; i < len(userEntity); i++ {
		for j := 0; j < len(cutiPengguna); j++ {
			if userEntity[i].ID == cutiPengguna[j].UserID {
				cutiPengguna[j].User = User(userEntity[i])
				cutiEntity = append(cutiEntity, PengunaToEntity(cutiPengguna[j]))
			}
		}
	}
	return cutiEntity, nil
}

// SelectAllKaryawan implements cuti.CutiDataInterface.
func (repo *CutiData) SelectAllKaryawan(idUser string) ([]cuti.CutiEntity, error) {
	//offset := (page - 1) * item
	var inputModel []Cuti
	tx := repo.db.Find(&inputModel)
	if tx.Error != nil {
		return nil, errors.New("error get all cuti karyawan")
	}

	dataUser, errUser := usernodejs.GetByIdUser(idUser)
	if errUser != nil {
		return nil, errUser
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
	return cutiEntity, nil
}

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
