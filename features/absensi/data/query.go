package data

import (
	"be_golang/klp3/features/absensi"
	apinodejs "be_golang/klp3/features/apiNodejs"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type absensiQuery struct {
	db          *gorm.DB
	externalAPI apinodejs.ExternalDataInterface
}

// SelectUserById implements absensi.AbsensiDataInterface
func (*absensiQuery) SelectUserById(idUser string) (absensi.PenggunaEntity, error) {
	data, err := usernodejs.GetByIdUser(idUser)
	if err != nil {
		return absensi.PenggunaEntity{}, err
	}
	dataUser := UserNodeJskePengguna(data)
	dataEntity := UserPenggunaToEntity(dataUser)
	return dataEntity, nil
}

// SelectById implements absensi.AbsensiDataInterface
func (repo *absensiQuery) SelectById(absensiID string) (absensi.AbsensiEntity, error) {
	var absensiData Absensi

	tx := repo.db.Where("id = ?", absensiID).First(&absensiData)
	if tx.Error != nil {
		log.Printf("Error read absensi: %s", tx.Error)
		return absensi.AbsensiEntity{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No rows affected when read absensi")
		return absensi.AbsensiEntity{}, errors.New("absensi not found")
	}
	//Mapping absensi to CorePabsensi
	coreAbsensi := ModelToEntity(absensiData)
	log.Println("Read absensi successfully")
	return coreAbsensi, nil
}

// GetUserByIDAPI implements absensi.AbsensiDataInterface
func (repo *absensiQuery) GetUserByIDAPI(idUser string) (apinodejs.Pengguna, error) {
	// Panggil metode GetUserByID dari externalAPI
	user, err := repo.externalAPI.GetUserByID(idUser)
	if err != nil {
		log.Printf("Error consume api user: %s", err.Error())
		return apinodejs.Pengguna{}, err
	}
	log.Println("consume api successfully")
	return user, nil
}

// Insert implements absensi.AbsensiDataInterface
func (repo *absensiQuery) Insert(input absensi.AbsensiEntity) error {
	idUser, errIdUser := helper.GenerateUUID()
	if errIdUser != nil {
		return errors.New("error generate uuid")
	}
	inputModel := EntityToModel(input)
	inputModel.ID = idUser
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("failed create absensi")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// Update implements absensi.AbsensiDataInterface
func (repo *absensiQuery) Update(input absensi.AbsensiEntity, idUser string, id string) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Absensi{}).Where("id=? and user_id=?", id, idUser).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("update absensi fail")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// SelectAll implements absensi.AbsensiDataInterface
func (repo *absensiQuery) SelectAll(param absensi.QueryParams) (int64, []absensi.AbsensiEntity, error) {
	var inputModel []Absensi
	var total_absensi int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		fmt.Println("offset", offset)
		if param.SearchName != "" {
			query = query.Where("nama_lengkap like ?", "%"+param.SearchName+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all absensi")
		}
		total_absensi = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	if param.SearchName != "" {
		query = query.Where("nama_lengkap like ?", "%"+param.SearchName+"%")
	}
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all absensi")
	}

	dataPengguna, errUser := usernodejs.GetAllUser()
	if errUser != nil {
		return 0, nil, errUser
	}
	var dataUser []User
	for _, value := range dataPengguna {
		dataUser = append(dataUser, PenggunaToUser(value))
	}
	var userEntity []absensi.UserEntity
	for _, value := range dataUser {
		userEntity = append(userEntity, UserToEntity(value))
	}
	fmt.Println("user entity", userEntity)
	var absensiPengguna []AbsensiPengguna
	for _, value := range inputModel {
		absensiPengguna = append(absensiPengguna, ModelToPengguna(value))
	}
	fmt.Println("reimb", absensiPengguna)
	var absensiEntity []absensi.AbsensiEntity
	for i := 0; i < len(userEntity); i++ {
		for j := 0; j < len(absensiPengguna); j++ {
			if userEntity[i].ID == absensiPengguna[j].UserID {
				absensiPengguna[j].User = User(userEntity[i])
				absensiEntity = append(absensiEntity, PenggunaToEntity(absensiPengguna[j]))
			}
		}
	}
	return total_absensi, absensiEntity, nil
}

// SelectAllKaryawan implements absensi.AbsensiDataInterface
func (repo *absensiQuery) SelectAllKaryawan(idUser string, param absensi.QueryParams) (int64, []absensi.AbsensiEntity, error) {
	var inputModel []Absensi
	var total_absensi int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		if param.SearchName != "" {
			query = query.Where("user_id=? and nama_lengkap like ?", idUser, "%"+param.SearchName+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all absensi")
		}
		total_absensi = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	if param.SearchName != "" {
		query = query.Where("user_id=? and nama_lengkap like ?", idUser, "%"+param.SearchName+"%")
	}
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all absensi karyawan")
	}

	dataUser, errUser := usernodejs.GetByIdUser(idUser)
	if errUser != nil {
		return 0, nil, errUser
	}
	pengguna := PenggunaToUser(dataUser)
	userEntity := UserToEntity(pengguna)

	var absensiPengguna []AbsensiPengguna
	for _, value := range inputModel {
		absensiPengguna = append(absensiPengguna, ModelToPengguna(value))
	}
	var absensiEntity []absensi.AbsensiEntity
	for _, value := range absensiPengguna {
		if value.UserID == userEntity.ID {
			value.User = User(userEntity)
			absensiEntity = append(absensiEntity, PenggunaToEntity(value))
		}
	}
	return total_absensi, absensiEntity, nil
}

func New(db *gorm.DB, externalAPI apinodejs.ExternalDataInterface) absensi.AbsensiDataInterface {
	return &absensiQuery{
		db:          db,
		externalAPI: externalAPI,
	}
}
