package doctor

type DoctorInfo struct {
	Id        int    `json:"-" gorm:"column:ID" gorm:"primaryKey"`
	DoctorId  string `json:"doctor_id" gorm:"column:DOCTOR_ID"`
	Title     string `json:"title" gorm:"column:TITLE"`
	FirstName string `json:"first_name" gorm:"column:FIRST_NAME"`
	LastName  string `json:"last_name" gorm:"column:LAST_NAME"`
}

type DoctorVisit struct {
	DoctorId string               `json:"doctor_id" `
	Schdule  []DoctorScheduleInfo `json:"doctor_schedule_info" `
}

type DoctorScheduleInfo struct {
	Id               string `json:"-" gorm:"column:ID" gorm:"primaryKey"`
	DoctorId         string `json:"-" gorm:"column:DOCTOR_ID"`
	ScheduleDay      string `json:"schedule_day" gorm:"column:SCHEDULE_DAY"`               //day code 0-6 >> sun to sat
	SchduleTimeStart string `json:"schedule_time_start" gorm:"column:SCHEDULE_TIME_START"` //HH:MM
	SchduleTimeEnd   string `json:"schedule_time_end" gorm:"column:SCHEDULE_TIME_END"`     //HH:MM
	ScheduletimeSlot int    `json:"schedule_time_slot" gorm:"column:SCHEDULE_TIME_SLOT"`   //minutes
	ActiveFlag       string `json:"active_flag" gorm:"column:ACTIVE_FLAG"`                 //minutes
}

type DoctorInScheduleInfo struct {
	Id       string `json:"-" gorm:"column:ID" gorm:"primaryKey"`
	DoctorId string `json:"doctor_id" gorm:"column:DOCTOR_ID"`
}
