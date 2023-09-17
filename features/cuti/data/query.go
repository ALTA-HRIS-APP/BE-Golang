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
		value.User = User(userEntity)
		cutiEntity = append(cutiEntity, PengunaToEntity(value))
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
