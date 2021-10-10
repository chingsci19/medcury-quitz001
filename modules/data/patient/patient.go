package patient

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBGetPatientInfoFromMobile(mobileNo string) PatientInfo {

	var patientInfo PatientInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return patientInfo
	}

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	patientInfoSearch := PatientInfo{MobileNo: mobileNo}
	db.Table("PATIENT").First(&patientInfo, patientInfoSearch)

	return patientInfo
}
