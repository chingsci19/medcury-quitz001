package appointment

import (
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func DBGetAppointment(doctorId string, startDate time.Time, endDate time.Time) []AppointmentInfo {

// 	var apnInfo []AppointmentInfo
// 	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

// 	if err != nil {
// 		return apnInfo
// 	}
// 	sqlDb, _ := db.DB()
// 	defer sqlDb.Close()

// 	if strings.TrimSpace(doctorId) == "" {
// 		db.Table("APPOINTMENT").Find(&apnInfo)
// 	} else {
// 		apnSearch := AppointmentInfo{DoctorId: doctorId}
// 		db.Table("APPOINTMENT").Where("APPOINTMENT_DATE BETWEEN  ? AND ? ", startDate, endDate).Find(&apnInfo, apnSearch)
// 	}

// 	return apnInfo
// }

func DBHaveAppointmentBySlotCode(slotCode string, status string) (bool, error) {

	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return true, err
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	var apnInfo []AppointmentInfo
	apnSearch := AppointmentInfo{SlotCode: slotCode, Status: status}
	var count int64
	db.Table("APPOINTMENT").Find(&apnInfo, apnSearch).Count(&count)
	if count <= 0 {
		return false, nil
	} else {
		return true, nil
	}

}

func DBGetAppointmentById(id int) (AppointmentInfo, error) {

	var apnInfo AppointmentInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return apnInfo, err
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	apnSearch := AppointmentInfo{Id: id}
	db.Table("APPOINTMENT").First(&apnInfo, apnSearch)

	return apnInfo, nil
}

func DBGetAppointmentBySlotCode(code string, status string) []AppointmentInfo {

	var apnInfo []AppointmentInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return apnInfo
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if strings.TrimSpace(status) == "" {
		status = "A"
	}

	apnSearch := AppointmentInfo{SlotCode: code, Status: strings.ToUpper(status)}
	db.Table("APPOINTMENT").Find(&apnInfo, apnSearch)

	return apnInfo
}

func DBSaveAppointment(appointment AppointmentInfo) (int64, error) {

	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return -1, err
	}

	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	tx := db.Table("APPOINTMENT").Save(&appointment)

	return tx.RowsAffected, nil
}

func DBGetAppointmentByDoctorId(docId int, status []string) ([]AppointmentInfo, error) {

	var apnInfo []AppointmentInfo
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.connection")), &gorm.Config{})

	if err != nil {
		return apnInfo, err
	}
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	apnSearch := AppointmentInfo{DoctorRefId: docId}
	if len(status) > 0 {
		db.Table("APPOINTMENT").Where("STATUS IN (?)", status).Find(&apnInfo, apnSearch)
	} else {
		db.Table("APPOINTMENT").Find(&apnInfo, apnSearch)
	}

	return apnInfo, nil
}
