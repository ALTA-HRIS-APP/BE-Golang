package service

import (
	"be_golang/klp3/features/target"
	"be_golang/klp3/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mocksTargetDataLayer := new(mocks.TargetData)
	t.Run("Success Create Target", func(t *testing.T) {
		userPembuat := target.PenggunaEntity{
			ID:          "54396f94-07b8-4450-8105-7c4472bf8701",
			NamaLengkap: "popol",
			Jabatan:     "manager",
		}
		userPenerima := target.PenggunaEntity{
			ID:          "27567353-9507-43d3-b08c-eea2c8c094fb",
			NamaLengkap: "vexana",
			Jabatan:     "karyawan",
		}
		mocksTargetDataLayer.On("GetUserByIDAPI", "54396f94-07b8-4450-8105-7c4472bf8701").Return(userPembuat, nil).Once()
		mocksTargetDataLayer.On("GetUserByIDAPI", "27567353-9507-43d3-b08c-eea2c8c094fb").Return(userPenerima, nil).Once()
		insertData := target.TargetEntity{
			KontenTarget:   "manajemen keuangan",
			Status:         "not completed",
			DevisiID:       "68a83bd8-a392-4877-b10f-f00251850cb8",
			UserIDPembuat:  "54396f94-07b8-4450-8105-7c4472bf8701",
			UserIDPenerima: "27567353-9507-43d3-b08c-eea2c8c094fb",
			DueDate:        "25-09-2023",
			Proofs:         "",
		}
		// Expectation on the mock
		mocksTargetDataLayer.On("Insert", insertData).Return(("1"), nil).Once()

		//object service layer dengan mock
		srv := New(mocksTargetDataLayer)
		createdTargetID, err := srv.Create(insertData)
		assert.Nil(t, err)
		assert.Equal(t, ("1"), createdTargetID)
		mocksTargetDataLayer.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	// Membuat objek mock TargetData
	mocksTargetDataLayer := new(mocks.TargetData)

	t.Run("Success: Karyawan viewing their own targets", func(t *testing.T) {
		userID := "54396f94-07b8-4450-8105-7c4472bf8701"
		param := target.QueryParam{
			Page:           int(1),
			LimitPerPage:   int(1),
			SearchKonten:   "menejemen",
			SearchStatus:   "not completed",
			ExistOtherPage: true,
		}

		// Membuat instance user dengan peran karyawan
		user := target.PenggunaEntity{
			ID:      userID,
			Jabatan: "karyawan",
		}

		// Mengatur bahwa pemanggilan metode GetUserByIDAPI akan mengembalikan user yang valid
		mocksTargetDataLayer.On("GetUserByIDAPI", userID).Return(user, nil).Once()

		// Mengatur bahwa pemanggilan metode SelectAllKaryawan akan mengembalikan data target yang valid
		// Sesuaikan dengan hasil yang diharapkan
		mocksTargetDataLayer.On("SelectAllKaryawan", userID, param).Return(int64(2), []target.TargetEntity{
			// Isi data target yang diharapkan
		}, nil).Once()

		srv := New(mocksTargetDataLayer)

		hasNextPage, data, err := srv.GetAll(userID, param)
		assert.Nil(t, err)
		assert.True(t, hasNextPage)
		// Memeriksa data yang dihasilkan sesuai dengan yang diharapkan
		// Sesuaikan dengan harapan Anda
		assert.Len(t, data, 2)

		mocksTargetDataLayer.AssertExpectations(t)
	})

	t.Run("Success: Non-Karyawan viewing all targets", func(t *testing.T) {
		// Kasus ini akan mirip dengan kasus sebelumnya, tetapi dengan peran yang berbeda
		// Pastikan untuk mengatur hasil pemanggilan metode SelectAll yang sesuai
	})

	t.Run("Success: No Next Page", func(t *testing.T) {
		// Kasus ini menguji ketika parameter ExistOtherPage = false
		// Anda dapat mengaturnya dengan parameter yang sesuai
		// dan memeriksa bahwa hasNextPage adalah false
	})

	// ... Pengujian lainnya ...
}
func TestGetById(t *testing.T) {
	// Membuat objek mock TargetData
	mocksTargetDataLayer := new(mocks.TargetData)

	t.Run("Success getting target details", func(t *testing.T) {
		targetID := "afd75070-9de2-4bef-be2c-cf60a63c719d"

		// Membuat data target yang diharapkan sebagai hasil dari pemanggilan Select
		expectedTarget := target.TargetEntity{
			ID:             "",
			KontenTarget:   "RAB",
			Status:         "completed",
			DevisiID:       "68a83bd8-a392-4877-b10f-f00251850cb8",
			UserIDPembuat:  "54396f94-07b8-4450-8105-7c4472bf8701",
			UserIDPenerima: "27567353-9507-43d3-b08c-eea2c8c094fb",
			DueDate:        "31-09-2023",
			Proofs:         "https://res.cloudinary.com/duklipjcj/image/upload/v1695210901/Screenshot%20%28173%29.png.png",
		}

		// Expectation pada mock
		mocksTargetDataLayer.On("Select", targetID).Return(expectedTarget, nil).Once()

		// Membuat instance targetService dengan mock
		srv := New(mocksTargetDataLayer)

		// Memanggil metode GetById
		result, err := srv.GetById(targetID, "afd75070-9de2-4bef-be2c-cf60a63c719d")

		// Memeriksa hasil
		assert.Nil(t, err)
		assert.Equal(t, expectedTarget, result)

		// Memeriksa ekspektasi pada mock
		mocksTargetDataLayer.AssertExpectations(t)
	})
	t.Run("Error Case: Get Target by ID", func(t *testing.T) {
		targetID := "afd75070-9de2-4bef-be2c-cf60a63c719d"

		// Mengatur bahwa pemanggilan metode Select akan mengembalikan error
		expectedErr := errors.New("Error getting target details")
		mocksTargetDataLayer.On("Select", targetID).Return(target.TargetEntity{}, expectedErr).Once()

		srv := New(mocksTargetDataLayer)

		result, err := srv.GetById(targetID, "afd75070-9de2-4bef-be2c-cf60a63c719d")

		// Memeriksa bahwa err adalah error yang diharapkan
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)

		// Memeriksa bahwa result adalah nilai nol atau sesuai dengan nilai default yang diharapkan
		assert.Equal(t, target.TargetEntity{}, result)

		// Memeriksa ekspektasi pada mock
		mocksTargetDataLayer.AssertExpectations(t)
	})
}

func TestUpdateTarget(t *testing.T) {}
func TestDeleteById(t *testing.T)   {}
