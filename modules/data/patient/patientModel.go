package patient

import "time"

type PatientInfo struct {
	Id          int       `json:"-" gorm:"column:ID" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"column:TITLE"`
	FirstName   string    `json:"first_name" gorm:"column:FIRST_NAME"`
	MiddleName  string    `json:"middle_name" gorm:"column:MIDDLE_NAME"`
	LastName    string    `json:"last_name" gorm:"column:LAST_NAME"`
	HN          string    `json:"hn" gorm:"column:HN"`
	MobileNo    string    `json:"mobile_no"  gorm:"column:MOBILE_NO" `
	Pin         string    `json:"pin"  gorm:"column:PIN_CODE"`
	Status      string    `json:"status"  gorm:"column:STATUS"`
	UpdatedBy   string    `json:"updated_by" gorm:"column:UPDATED_BY"`
	UpdatedDate time.Time `json:"updated_date" gorm:"column:UPDATED_DATE"`
}
