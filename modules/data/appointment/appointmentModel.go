package appointment

import "time"

type AppointmentInfo struct {
	Id                   int       `json:"id" gorm:"column:ID" gorm:"primary_key;auto_increment;not_null"`
	DoctorRefId          int       `json:"doctor_ref_id" gorm:"column:DOCTOR_REF_ID"`
	DoctorId             string    `json:"doctor_id" gorm:"column:DOCTOR_ID"`
	DoctorName           string    `json:"doctor_name" gorm:"column:DOCTOR_NAME"`
	PatientId            int       `json:"patient_id" gorm:"column:PATIENT_ID"`
	HN                   string    `json:"hn" gorm:"column:HN"`
	PatientName          string    `json:"patient_name" gorm:"column:PATIENT_NAME"`
	AppointmentDate      time.Time `json:"appointment_date" gorm:"column:APPOINTMENT_DATE"`
	AppointmentTimeStart time.Time `json:"appointment_time_start" gorm:"column:APPOINTMENT_START"`
	AppointmentTimeEnd   time.Time `json:"appointment_time_end" gorm:"column:APPOINTMENT_END"`
	Status               string    `json:"status" gorm:"column:STATUS"`
	UpdatedBy            string    `json:"updated_by" gorm:"column:UPDATED_BY"`
	UpdatedDate          time.Time `json:"updated_date" gorm:"column:UPDATED_DATE"`
	Remark               *string   `json:"remark" gorm:"column:REMARK"`
	SlotCode             string    `json:"slot_code" gorm:"column:SLOT_CODE"`
	PatientMobileNo      string    `json:"patient_mobile_no" gorm:"column:PATIENT_MOBILE_NO"`
}
