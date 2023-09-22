package service

// import (
// 	"be_golang/klp3/features/reimbusment"
// 	"be_golang/klp3/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestDelete(t *testing.T) {
// 	repo:=new(mocks.ReimbursementData)

// 	t.Run("delete success reimbursement",func(t *testing.T){
// 		repo.On("Delete","7198292739").Return(nil).Once()
// 		srv:=New(repo)
// 		err:=srv.Delete("7198292739")
// 		assert.Nil(t,err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("delete failed",func(t *testing.T){
// 		repo.On("Delete","7198292739").Return(errors.New("failed delete reimbursement")).Once()
// 		srv:=New(repo)
// 		err:=srv.Delete("7198292739")
// 		assert.NotNil(t,err)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetById(t *testing.T) {
// 	repo:=new(mocks.ReimbursementData)
// 	returnData:=reimbusment.ReimbursementEntity{
// 		ID: "7198292739",Description: "alat tulis kantor",Nominal: int(10000),UserID: "910394029",
// 	}
// 	returnUser:=reimbusment.PenggunaEntity{
// 		ID: "71982283739",NamaLengkap: "santi",
// 	}
// 	t.Run("success get by id",func(t *testing.T){
// 		repo.On("SelectById","7198292739").Return(returnData,nil).Once()
// 		repo.On("SelectUserById",returnData.UserID).Return(returnUser,nil).Once()
// 		srv:=New(repo)
// 		response,err:=srv.GetReimbusherById("7198292739")
// 		assert.Nil(t,err)
// 		assert.Equal(t,returnData.ID,response.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("failed select reimbursement",func(t *testing.T){
// 		repo.On("SelectById","7198292739").Return(reimbusment.ReimbursementEntity{},errors.New("failed reimbursement")).Once()
// 		srv:=New(repo)
// 		response,err:=srv.GetReimbusherById("7198292739")
// 		assert.NotNil(t,err)
// 		assert.Equal(t,reimbusment.ReimbursementEntity{},response)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("failed select user",func(t *testing.T){
// 		repo.On("SelectById","7198292739").Return(returnData,nil).Once()
// 		repo.On("SelectUserById",returnData.UserID).Return(reimbusment.PenggunaEntity{},errors.New("failed get user")).Once()
// 		srv:=New(repo)
// 		response,err:=srv.GetReimbusherById("7198292739")
// 		assert.Nil(t,err)
// 		assert.Equal(t,reimbusment.ReimbursementEntity{},response)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetAll(t *testing.T) {
// 	repo:=new(mocks.ReimbursementData)
// 	returnData:=[]reimbusment.ReimbursementEntity{
// 		{
// 		ID:              "19283748",
// 		Description:     "bola voli",
// 		Status:          "pending",
// 		BatasanReimburs: 5000000,
// 		Nominal:         20000,
// 		Tipe:            "rekreasi",
// 		Date:            "08 Juni 2023",
// 		Persetujuan:     "Done",
// 		UrlBukti:        "ikamska.jpg",
// 		UserID:          "9182930",
// },{
// 			ID:              "18293034",
// 			Description:     "bola voli",
// 			Status:          "pending",
// 			BatasanReimburs: 5000000,
// 			Nominal:         30000,
// 			Tipe:            "rekreasi",
// 			Date:            "08 Juni 2023",
// 			Persetujuan:     "Done",
// 			UrlBukti:        "ikamska.jpg",
// 			UserID:          "9182930",
// 			},}
// 	returnUserKaryawan:=reimbusment.PenggunaEntity{
// 		ID: "9182930",
// 		NamaLengkap: "sandi",
// 		Jabatan: "karyawan",
// 	}
// 	returnUserNotKaryawan:=reimbusment.PenggunaEntity{
// 		ID: "9182930",
// 		NamaLengkap: "sandi",
// 		Jabatan: "manager",
// 	}
// 	inputParam:=reimbusment.QueryParams{
// 		Page:             int(1),
// 		ItemsPerPage:     int(1),
// 		SearchName:       "bola",
// 		IsClassDashboard: true,
// 	}

// 	t.Run("success get all karyawan",func(t *testing.T) {
// 		repo.On("SelectUserById","9182930").Return(returnUserKaryawan,nil).Once()
// 		repo.On("SelectAllKaryawan","9182930",inputParam).Return(int64(56),returnData,nil).Once()
// 		srv:=New(repo)
// 		_,response,err:=srv.Get("9182930",inputParam)
// 		assert.Nil(t,err)
// 		assert.Equal(t,returnData[0].ID,response[0].ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("success get all not karyawan",func(t *testing.T) {
// 		repo.On("SelectUserById","9182930").Return(returnUserNotKaryawan,nil).Once()
// 		repo.On("SelectAll",inputParam).Return(int64(56),returnData,nil).Once()
// 		srv:=New(repo)
// 		_,response,err:=srv.Get("9182930",inputParam)
// 		assert.Nil(t,err)
// 		assert.Equal(t,returnData[0].ID,response[0].ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("failed get user",func(t *testing.T) {
// 		repo.On("SelectUserById","9182930").Return(reimbusment.PenggunaEntity{},errors.New("error get user")).Once()
// 		srv:=New(repo)
// 		_,response,err:=srv.Get("9182930",inputParam)
// 		assert.NotNil(t,err)
// 		assert.Nil(t,response)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("failed get all karyawan",func(t *testing.T) {
// 		repo.On("SelectUserById","9182930").Return(returnUserKaryawan,nil).Once()
// 		repo.On("SelectAllKaryawan","9182930",inputParam).Return(int64(0),nil,errors.New("error get all karyawan")).Once()
// 		srv:=New(repo)
// 		_,response,err:=srv.Get("9182930",inputParam)
// 		assert.NotNil(t,err)
// 		assert.Nil(t,response)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("failed get all not karyawan",func(t *testing.T) {
// 		repo.On("SelectUserById","9182930").Return(returnUserNotKaryawan,nil).Once()
// 		repo.On("SelectAll",inputParam).Return(int64(0),nil,errors.New("failed get reimbursement")).Once()
// 		srv:=New(repo)
// 		_,response,err:=srv.Get("9182930",inputParam)
// 		assert.NotNil(t,err)
// 		assert.Nil(t,response)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestCreate(t *testing.T) {
// 	repo:=new(mocks.ReimbursementData)
// 	inputData:= reimbusment.ReimbursementEntity{
// 		ID:              "91039404048",
// 		Description:     "sewa gedung",
// 		Nominal:         2000000,
// 		Tipe:            "sewa",
// 		Date:            "21 September 2023",
// 		UrlBukti:        "https://iqjkajwuwjs.jpg",
// 		UserID:          "81729201829",
// 	}
// 	t.Run("success create",func(t *testing.T) {
// 		repo.On("Insert",inputData).Return(nil).Once()
// 		srv:=New(repo)
// 		err:=srv.Add(inputData)
// 		assert.Nil(t,err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("failed create",func(t *testing.T) {
// 		repo.On("Insert",inputData).Return(errors.New("error create data")).Once()
// 		srv:=New(repo)
// 		err:=srv.Add(inputData)
// 		assert.NotNil(t,err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("validate create",func(t *testing.T) {
// 		inputData:= reimbusment.ReimbursementEntity{
// 			ID:              "91039404048",
// 			Description:     "sewa gedung",
// 			Nominal:         2000000,
// 			Tipe:            "sewa",
// 			Date:            "21 September 2023",
// 		}
// 			srv:=New(repo)
// 			err:=srv.Add(inputData)
// 			assert.NotNil(t,err)
// 			repo.AssertExpectations(t)
// 	})

// 	t.Run("validate batasan reimbursement",func(t *testing.T) {
// 		inputDataBatasan:= reimbusment.ReimbursementEntity{
// 			ID:              "91039404048",
// 			Description:     "sewa gedung",
// 			Nominal:         700000000000,
// 			Tipe:            "sewa",
// 			Date:            "21 September 2023",
// 		}
// 			srv:=New(repo)
// 			err:=srv.Add(inputDataBatasan)
// 			assert.NotNil(t,err)
// 			repo.AssertExpectations(t)
// 	})

// }

// func TestEdit(t *testing.T) {
// 	repo:=new(mocks.ReimbursementData)
// 	inputData:= reimbusment.ReimbursementEntity{
// 		Description:     "sewa gedung",
// 		Date:            "21 September 2023",
// 		UrlBukti:        "https://iqjkajwuwjs.jpg",
// 		UserID:          "81729201829",
// 	}

// 	returnUserKaryawan:=reimbusment.PenggunaEntity{
// 		ID:          "9283920282",
// 		NamaLengkap: "rudi",
// 		Jabatan:     "karyawan",
// 		Devisi: "produksi",
// 	}
// 	t.Run("success edit karyawan",func(t *testing.T) {
// 		repo.On("SelectUserById","81729201829").Return(returnUserKaryawan,nil).Once()
// 		repo.On("SelectById","91039404048").Return(inputData,nil).Once()
// 		repo.On("SelectUserById","81729201829").Return(returnUserKaryawan,nil).Once()
// 		repo.On("UpdateKaryawan",inputData,"91039404048").Return(nil).Once()
// 		srv:=New(repo)
// 		err:=srv.Edit(inputData,"91039404048","81729201829")
// 		assert.Nil(t,err)
// 		repo.AssertExpectations(t)
// 	})

// }