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
		assert.Equal(t, "1", createdTargetID)
		mocksTargetDataLayer.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	// Membuat objek mock TargetData
	mocksTargetDataLayer := new(mocks.TargetData)

	t.Run("Success getting target details", func(t *testing.T) {
		targetID := "afd75070-9de2-4bef-be2c-cf60a63c719d"

		// Membuat data target yang diharapkan sebagai hasil dari pemanggilan Select
		expectedTarget := target.TargetEntity{
			ID:             "afd75070-9de2-4bef-be2c-cf60a63c719d",
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

func TestDeleteById(t *testing.T) {
	// Membuat instance mock untuk TargetData
	mocksTargetDataLayer := new(mocks.TargetData)

	// Membuat instance targetService dengan mock
	srv := New(mocksTargetDataLayer)

	// ID target yang akan dihapus dan ID pengguna yang melakukan penghapusan
	targetID := "1"
	userID := "54396f94-07b8-4450-8105-7c4472bf8701" // Ganti dengan userID yang sesuai dengan pengujian ini

	// Data pengguna pembuat target
	userPembuat := target.PenggunaEntity{
		ID:          "54396f94-07b8-4450-8105-7c4472bf8701", // Menggunakan userID yang sesuai
		NamaLengkap: "popol",
		Jabatan:     "manager",
	}

	// Data target yang akan dihapus (sesuaikan dengan yang sesuai dengan pengujian ini)
	targetToDelete := target.TargetEntity{
		UserIDPembuat:  "54396f94-07b8-4450-8105-7c4472bf8701", // Menggunakan userID yang sesuai
		UserIDPenerima: "27567353-9507-43d3-b08c-eea2c8c094fb",
	}

	t.Run("Success Delete Target", func(t *testing.T) {
		// Expectation: Mock akan memanggil GetUserByIDAPI dengan userID yang sesuai
		mocksTargetDataLayer.On("GetUserByIDAPI", "54396f94-07b8-4450-8105-7c4472bf8701").Return(userPembuat, nil).Once()

		// Expectation: Mock akan memanggil Select untuk mendapatkan data target yang akan dihapus
		mocksTargetDataLayer.On("Select", targetID).Return(targetToDelete, nil).Once()

		// Expectation: Mock akan memanggil Delete untuk menghapus target
		mocksTargetDataLayer.On("Delete", targetID).Return(nil).Once()

		// Menggunakan targetService untuk menghapus target
		err := srv.DeleteById(targetID, userID)

		// Memastikan tidak ada error yang terjadi
		assert.Nil(t, err)

		// Memastikan semua ekspektasi terpenuhi
		mocksTargetDataLayer.AssertExpectations(t)
	})
}
