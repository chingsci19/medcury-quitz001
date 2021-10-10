package doctor

import (
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBGetDoctorInfo(doctorId string) (DoctorInfo, error) {

	var docInfo DoctorInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return docInfo, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return docInfo, err
	}
	defer sqlDb.Close()

	docSearch := DoctorInfo{DoctorId: doctorId}
	db.Table("DOCTOR").First(&docInfo, docSearch)

	return docInfo, nil
}

func DBGetDoctorScheduleInfo(doctorId string, scheduleDay []int) []DoctorScheduleInfo {

	var docVisit []DoctorScheduleInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return docVisit
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	if strings.TrimSpace(doctorId) == "" {
		docSearch := DoctorScheduleInfo{ActiveFlag: "Y"}
		db.Table("DOCTOR_SCHEDULE").Order("DOCTOR_ID, SCHEDULE_DAY").Where("SCHEDULE_DAY IN (?)", scheduleDay).Find(&docVisit, docSearch)
	} else {
		docSearch := DoctorScheduleInfo{DoctorId: doctorId, ActiveFlag: "Y"}
		db.Table("DOCTOR_SCHEDULE").Order("DOCTOR_ID, SCHEDULE_DAY").Where("SCHEDULE_DAY IN (?)", scheduleDay).Find(&docVisit, docSearch)
	}

	return docVisit
}

func DBGetDoctorInScheduleInfo(doctorId string, scheduleDay []int) []DoctorInScheduleInfo {

	var docVisit []DoctorInScheduleInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return docVisit
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	if strings.TrimSpace(doctorId) == "" {
		docSearch := DoctorScheduleInfo{ActiveFlag: "Y"}
		db.Table("DOCTOR_SCHEDULE").Distinct("DOCTOR_ID").Order("DOCTOR_ID, SCHEDULE_DAY").Where("SCHEDULE_DAY IN (?)", scheduleDay).Find(&docVisit, docSearch)
	} else {
		docSearch := DoctorScheduleInfo{DoctorId: doctorId, ActiveFlag: "Y"}
		db.Table("DOCTOR_SCHEDULE").Distinct("DOCTOR_ID").Order("DOCTOR_ID, SCHEDULE_DAY").Where("SCHEDULE_DAY IN (?)", scheduleDay).Find(&docVisit, docSearch)
	}

	return docVisit
}
