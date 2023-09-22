package database

import (
	"be_golang/klp3/app/config"
	absensi "be_golang/klp3/features/absensi/data"
	cuti "be_golang/klp3/features/cuti/data"
	reimbusment "be_golang/klp3/features/reimbusment/data"
	target "be_golang/klp3/features/target/data"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&absensi.Absensi{}, &cuti.Cuti{}, &reimbusment.Reimbursement{}, target.Target{})
}
