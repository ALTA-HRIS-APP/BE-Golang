package service

import (
	"be_golang/klp3/features/target"
	"be_golang/klp3/mocks"
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
			Due_Date:       "25-09-2023",
			Proofs:         "https://res.cloudinary.com/duklipjcj/image/upload/v1695210901/Screenshot%20%28173%29.png.png",
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
