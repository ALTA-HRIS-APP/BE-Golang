package data

import (
	"be_golang/klp3/features/absensi"
	usernodejs "be_golang/klp3/features/userNodejs"
	"be_golang/klp3/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type absensiQuery struct {
	db *gorm.DB
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
		if param.SearchName != "" {
			query = query.Where("description like ?", "%"+param.SearchName+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all absensi")
		}
		total_absensi = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	if param.SearchName != "" {
		query = query.Where("description like ?", "%"+param.SearchName+"%")
	}
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all absensi")
	}

	// Tambahkan kode untuk mendapatkan tanggal sekarang
	now := time.Now()

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
	var absensiPengguna []AbsensiPengguna
	for _, value := range inputModel {
		absensiPengguna = append(absensiPengguna, ModelToPengguna(value))
	}
	var absensiEntity []absensi.AbsensiEntity
	for i := 0; i < len(userEntity); i++ {
		for j := 0; j < len(absensiPengguna); j++ {
			if userEntity[i].ID == absensiPengguna[j].UserID {
				absensiPengguna[j].User = User(userEntity[i])

				// Setel tanggal sekarang ke absensiEntity
				absensiPengguna[j].TanggalSekarang = now

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
			query = query.Where("user_id=? and description like ?", idUser, "%"+param.SearchName+"%")
		}
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get all absensi")
		}
		total_absensi = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	if param.SearchName != "" {
		query = query.Where("user_id=? and description like ?", idUser, "%"+param.SearchName+"%")
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

// SelectById implements absensi.AbsensiDataInterface
func (repo *absensiQuery) SelectById(id string) (absensi.AbsensiEntity, error) {
	var inputModel Absensi
	tx := repo.db.Where("id=?", id).First(&inputModel)
	if tx.Error != nil {
		return absensi.AbsensiEntity{}, errors.New("error get batasan reimbursment")
	}
	output := ModelToEntity(inputModel)
	return output, nil
}

func New(db *gorm.DB) absensi.AbsensiDataInterface {
	return &absensiQuery{
		db: db,
	}
}
