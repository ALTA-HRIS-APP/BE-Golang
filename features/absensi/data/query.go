package data

import (
	"be_golang/klp3/features/absensi"

	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type absensiQuery struct {
	db          *gorm.DB
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
func (repo *absensiQuery) GetUserByIDAPI(idUser string) (usernodejs.Pengguna, error) {
	// Panggil metode GetUserByID dari externalAPI
	user, err := usernodejs.GetByIdUser(idUser)
	if err != nil {
		log.Printf("Error consume api user: %s", err.Error())
		return usernodejs.Pengguna{}, err
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
func (repo *absensiQuery) SelectAll(token string, param absensi.QueryParams) (int64, []absensi.AbsensiEntity, error) {
	var inputModel []Absensi
	var total_absensi int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		fmt.Println("offset", offset)

		// Tambahkan kondisi untuk filter berdasarkan tanggal_sekarang
		if param.SerachTanggal != "" {
			// Parsing tanggal_sekarang dari format "2006-01-02"
			parsedDate, err := time.Parse("2006-01-02", param.SerachTanggal)
			if err != nil {
				return 0, nil, errors.New("failed to parse tanggal_sekarang")
			}
			// Filter data berdasarkan tanggal_sekarang
			query = query.Where("created_at >= ? AND created_at <= ?", parsedDate, parsedDate.Add(24*time.Hour))
		}

		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all absensi")
		}
		total_absensi = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}

	// Tambahkan kondisi untuk filter berdasarkan tanggal
	if param.SerachTanggal != "" && !param.IsClassDashboard {
		// Parsing tanggal_sekarang dari format "2006-01-02"
		parsedDate, err := time.Parse("2006-01-02", param.SerachTanggal)
		if err != nil {
			return 0, nil, errors.New("failed to parse tanggal_sekarang")
		}
		// Filter data berdasarkan tanggal_sekarang
		query = query.Where("created_at >= ? AND created_at <= ?", parsedDate, parsedDate.Add(24*time.Hour))
	}

	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all absensi")
	}
	dataPengguna, errUser := usernodejs.GetAllUser(token)
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

	var absensiPengguna []AbsensiPengguna
	for _, value := range inputModel {
		absensiPengguna = append(absensiPengguna, ModelToPengguna(value))
	}

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

		// Tambahkan kondisi untuk filter berdasarkan tanggal
		if param.SerachTanggal != "" {
			// Parsing tanggal dari format "2006-01-02"
			parsedDate, err := time.Parse("2006-01-02", param.SerachTanggal)
			if err != nil {
				return 0, nil, errors.New("failed to parse tanggal")
			}
			// Filter data berdasarkan tanggal
			query = query.Where("user_id=? AND DATE(created_at) = ?", idUser, parsedDate)
		}

		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all absensi")
		}
		total_absensi = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	} else {
		// Tambahkan kondisi untuk filter berdasarkan tanggal
		if param.SerachTanggal != "" {
			// Parsing tanggal dari format "2006-01-02"
			parsedDate, err := time.Parse("2006-01-02", param.SerachTanggal)
			if err != nil {
				return 0, nil, errors.New("failed to parse tanggal")
			}
			// Filter data berdasarkan tanggal
			query = query.Where("user_id=? AND DATE(created_at) = ?", idUser, parsedDate)
		}

		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("error get all absensi karyawan")
		}
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
	log.Println("select all karyawan", absensiEntity)
	return total_absensi, absensiEntity, nil
}

func New(db *gorm.DB) absensi.AbsensiDataInterface {
	return &absensiQuery{
		db:          db,

	}
}
