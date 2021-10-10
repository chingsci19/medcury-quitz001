package appointment

import (
	"medcury/quitz001/modules/data/doctor"
	"time"
)

type AppointmentListReq struct {
	StartDate string `json:"start_date" ` //YYYY-MM-DD
	EndDate   string `json:"end_date" `   //YYYY-MM-DD
}

type AppointmentListRes struct {
	AppointmentListInfo []AppointmentListInfo `json:"appointment_list_info" `
}

type AppointmentListInfo struct {
	DoctorInfo           doctor.DoctorInfo           `json:"doctor_info" `
	Schdule              []doctor.DoctorScheduleInfo `json:"doctor_schedule_info" `
	AvailableAppointment []AvailableAppointment      `json:"available_appointment" `
}

type AvailableAppointment struct {
	SlotDay         string `json:"slot_day" `
	SlotPeriodStart string `json:"slot_period_start" `
	SlotPeriodEnd   string `json:"slot_period_end" `
	SlotPeriodText  string `json:"slot_period_text" `
	IsAvaliable     bool   `json:"avaliable" `
	SlotCode        string `json:"slot_code"`
}

type MakeAppointmentReq struct {
	MobileNo string `json:"mobile_no" `
	Pin      string `json:"pin" `
	SlotCode string `json:"slot_code" `
}

type DoctorAppointmentReportListRes struct {
	DoctorInfo  doctor.DoctorInfo           `json:"doctor_info" `
	Schdule     []doctor.DoctorScheduleInfo `json:"doctor_schedule_info" `
	Appointment []AppointmentInfoItem       `json:"appointment" `
}

type AppointmentInfoItem struct {
	HN                   string    `json:"hn"`
	PatientName          string    `json:"patient_name"`
	AppointmentDate      time.Time `json:"appointment_date"`
	AppointmentTimeStart time.Time `json:"appointment_time_start"`
	AppointmentTimeEnd   time.Time `json:"appointment_time_end"`
	Status               string    `json:"status" `
	UpdatedBy            string    `json:"updated_by"`
	UpdatedDate          time.Time `json:"updated_date" `
	Remark               *string   `json:"remark" `
	SlotCode             string    `json:"slot_code"`
	PatientMobileNo      string    `json:"patient_mobile_no" `
}
