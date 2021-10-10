package appointment

import (
	"medcury/quitz001/services/base"
	"regexp"
	"strconv"
	"strings"
	"time"

	"medcury/quitz001/modules/data/appointment"
	"medcury/quitz001/modules/data/doctor"
	"medcury/quitz001/modules/data/patient"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const location = "Asia/Bangkok"

func Route(group fiber.Router) {
	group.Post("/list", func(c *fiber.Ctx) error {
		vResp := AppointmentList(c)
		log.Debug().Msgf("vResp : %v", vResp)
		return c.Status(vResp.HttpStatus).JSON(vResp)
	})

	group.Put("", func(c *fiber.Ctx) error {
		vResp := MakeAppointment(c)
		log.Debug().Msgf("vResp : %v", vResp)
		return c.Status(vResp.HttpStatus).JSON(vResp)
	})

	group.Delete("", func(c *fiber.Ctx) error {
		vResp := CancelAppointment(c)
		log.Debug().Msgf("vResp : %v", vResp)
		return c.Status(vResp.HttpStatus).JSON(vResp)
	})

	group.Get("/doctor", func(c *fiber.Ctx) error {
		vResp := GetDoctorAppointmentReport(c)
		log.Debug().Msgf("vResp : %v", vResp)
		return c.Status(vResp.HttpStatus).JSON(vResp)
	})

}

func AppointmentList(c *fiber.Ctx) base.Response {
	logData := make(map[string]interface{})
	logData["tag"] = "appointment.AppointmentList"
	bResponse := base.Response{}

	appointmentReq := AppointmentListReq{}
	err := c.BodyParser(&appointmentReq)

	if err != nil {
		// log.Err(err).Fields(logData).Caller().Msg(err.Error())
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}

	//validate and convert string to date
	startDate, err := time.Parse("2006-01-02", appointmentReq.StartDate)
	if err != nil {
		log.Error().Fields(logData).Caller().Msg(err.Error())
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}

	endDate, err := time.Parse("2006-01-02", appointmentReq.EndDate)
	if err != nil {
		log.Error().Fields(logData).Caller().Msg(err.Error())
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}

	if endDate.Before(time.Now()) || startDate.Before(time.Now()) {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-003"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}

	if endDate.Before(startDate) {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-002"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	// log.Debug().Msgf("%i", startDate.Weekday())

	scheduleDay := []int{0, 1, 2, 3, 4, 5, 6} //all days
	// scheduleDay = append(scheduleDay, int(startDate.Weekday()))
	diffDate := endDate.Sub(startDate).Hours() / 24
	if diffDate < 7 {
		scheduleDay = []int{} //clear days
		for i := 0; i <= int(diffDate); i++ {
			day := int(startDate.Weekday()) + i
			if day <= 6 {
				scheduleDay = append(scheduleDay, day)
			} else {
				scheduleDay = append(scheduleDay, day-7)
			}

		}
	}
	log.Debug().Msgf("%v", scheduleDay)
	//
	// find all doctor schedule by day of week
	// docSchedule := doctor.DBGetDoctorScheduleInfo("", scheduleDay)
	// log.Debug().Msgf("%v", docSchedule)

	docInSchedule := doctor.DBGetDoctorInScheduleInfo("", scheduleDay)
	log.Debug().Msgf("%v", docInSchedule)

	if len(docInSchedule) > 0 {
		var appointmentListRes AppointmentListRes
		for _, v := range docInSchedule {
			appointmentListInfo := AppointmentListInfo{}
			docInfo, err := doctor.DBGetDoctorInfo(v.DoctorId)
			if err != nil {
				bResponse.HttpStatus = 500
				bResponse.StatusCode = "APP-500"
				bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
				bResponse.Data = err.Error()
				return bResponse
			}
			docSchedule := doctor.DBGetDoctorScheduleInfo(v.DoctorId, scheduleDay)
			docAllSchedule := doctor.DBGetDoctorScheduleInfo(v.DoctorId, []int{0, 1, 2, 3, 4, 5, 6})
			// doctorAppointment := appointment.DBGetAppointment(v.DoctorId, startDate, endDate)
			log.Debug().Msgf("%v", docInfo)
			availableAppointment := CreateDoctorAvailableAppointment(docSchedule, startDate, endDate)
			log.Debug().Msgf("%v", availableAppointment)
			appointmentListInfo.DoctorInfo = docInfo
			appointmentListInfo.Schdule = docAllSchedule
			appointmentListInfo.AvailableAppointment = availableAppointment
			appointmentListRes.AppointmentListInfo = append(appointmentListRes.AppointmentListInfo, appointmentListInfo)
		}
		bResponse.Data = appointmentListRes
		bResponse.HttpStatus = 200
		bResponse.StatusCode = "APP-200"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)

	}
	// docInfos := doctor.DBGetDoctorInfo("001")

	// log.Debug().Msgf("%v", docInfos)

	return bResponse
}

func CreateDoctorAvailableAppointment(docSchedule []doctor.DoctorScheduleInfo, startDate time.Time, endDate time.Time) []AvailableAppointment {
	var availableAppointment []AvailableAppointment
	diffDate := endDate.Sub(startDate).Hours() / 24
	// log.Debug().Msgf("doctorAppointment : %v", doctorAppointment)

	for i := 0; i <= int(diffDate); i++ { //create day slot
		day := startDate.AddDate(0, 0, i)
		for _, v := range docSchedule {
			scheduleDay, _ := strconv.Atoi(v.ScheduleDay)
			if scheduleDay == int(day.Weekday()) { //ต้องเป็นวันที่ออกตรวจ
				t1, _ := time.Parse("15:04", v.SchduleTimeStart)
				timeStart := day.Add(time.Hour*time.Duration(t1.Hour()) + time.Minute*time.Duration(t1.Minute()))
				t1, _ = time.Parse("15:04", v.SchduleTimeEnd)
				timeEnd := day.Add(time.Hour*time.Duration(t1.Hour()) + time.Minute*time.Duration(t1.Minute()))

				for timeEnd.After(timeStart) {
					timeStartNext := timeStart.Add(time.Minute * time.Duration(v.ScheduletimeSlot))
					availableAppointmentItem := AvailableAppointment{SlotDay: day.Format("2006-01-02"),
						SlotPeriodStart: timeStart.Format("2006-01-02 15:04"),
						SlotPeriodEnd:   timeStartNext.Format("2006-01-02 15:04"),
						SlotPeriodText:  timeStart.Format("15:04") + "-" + timeStartNext.Format("15:04"),
						SlotCode:        v.DoctorId + "-" + timeStart.Format("200601021504") + "-" + timeStartNext.Format("1504"),
						IsAvaliable:     true}
					have, _ := appointment.DBHaveAppointmentBySlotCode(availableAppointmentItem.SlotCode, "A")
					if have {
						availableAppointmentItem.IsAvaliable = false
					}
					availableAppointment = append(availableAppointment, availableAppointmentItem)
					timeStart = timeStartNext
				}

			}
		}

	}
	return availableAppointment
}

func MakeAppointment(c *fiber.Ctx) base.Response {
	logData := make(map[string]interface{})
	logData["tag"] = "appointment.MakeAppointment"
	bResponse := base.Response{}

	appointmentReq := MakeAppointmentReq{}
	err := c.BodyParser(&appointmentReq)

	if err != nil {
		// log.Err(err).Fields(logData).Caller().Msg(err.Error())
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}
	bkk, err := time.LoadLocation(location)
	if err != nil {
		// log.Err(err).Fields(logData).Caller().Msg(err.Error())
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}
	if strings.TrimSpace(appointmentReq.MobileNo) == "" { // invalid mobile no
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}

	pinRegexp := regexp.MustCompile(`\d{6}`)
	if !pinRegexp.MatchString(appointmentReq.Pin) { // invalid pin format
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}

	patientInfo := patient.DBGetPatientInfoFromMobile(appointmentReq.MobileNo)
	if patientInfo.Id == 0 {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-404-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	if patientInfo.Pin != appointmentReq.Pin {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	//verify timeslot
	slotData := strings.Split(appointmentReq.SlotCode, "-")
	if len(slotData) != 3 { //ผิด format
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	if len(strings.TrimSpace(slotData[0])) == 0 { //ผิด format หมอ
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	if len(strings.TrimSpace(slotData[1])) != 12 { //ผิด format เวลาเริ่ม
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	if len(strings.TrimSpace(slotData[2])) != 4 { //ผิด format เวลาสิ้นสุด
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	}
	//ผ่านทั้งหมด เตรียมทำนัดหมาย
	appointmentInfos := appointment.DBGetAppointmentBySlotCode(appointmentReq.SlotCode, "A")
	if len(appointmentInfos) > 0 { //มีนัดหมายซ้ำ
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-004"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
		return bResponse
	} else { //ไม่มี
		doctorInfo, err := doctor.DBGetDoctorInfo(strings.TrimSpace(slotData[0]))
		if err != nil {
			bResponse.HttpStatus = 500
			bResponse.StatusCode = "APP-500"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = err.Error()
			return bResponse
		}
		if doctorInfo.Id <= 0 {
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-404-002"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = nil
			log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
			return bResponse
		}

		startTime, err := time.ParseInLocation("200601021504", slotData[1], bkk)
		if err != nil { //wrong date format
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-400-001"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = err.Error()
			return bResponse
		}
		endTime, err := time.ParseInLocation("200601021504", slotData[1][0:8]+slotData[2], bkk)
		if err != nil { //wrong date format
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-400-001"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = err.Error()
			return bResponse
		}

		day, _ := time.ParseInLocation("20060102", slotData[1][0:8], bkk)

		if endTime.Before(startTime) || endTime.Equal(startTime) { //เวลาสิ้นสุดต้องมากกว่าเวลาเริ่ม
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-400-001"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = nil
			log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
			return bResponse
		}
		dayOfWeek := []int{int(startTime.Weekday())}
		doctorScheduleInfo := doctor.DBGetDoctorScheduleInfo(doctorInfo.DoctorId, dayOfWeek)
		if len(doctorScheduleInfo) == 0 {
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-404-003"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = nil
			log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
			return bResponse
		}

		diffTime := endTime.Sub(startTime).Minutes()
		log.Debug().Msgf("%v", doctorScheduleInfo)
		if diffTime != float64(doctorScheduleInfo[0].ScheduletimeSlot) { //wrong slot duration
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-400-001"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = nil
			log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
			return bResponse
		}
		if startTime.Before(time.Now()) {
			bResponse.HttpStatus = 400
			bResponse.StatusCode = "APP-400-005"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = nil
			log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
			return bResponse
		}
		for _, v := range doctorScheduleInfo {
			t1, _ := time.Parse("15:04", v.SchduleTimeStart)
			timeStart := day.Add(time.Hour*time.Duration(t1.Hour()) + time.Minute*time.Duration(t1.Minute()))
			t1, _ = time.Parse("15:04", v.SchduleTimeEnd)
			timeEnd := day.Add(time.Hour*time.Duration(t1.Hour()) + time.Minute*time.Duration(t1.Minute()))
			log.Debug().Msgf("%v", startTime.Equal(timeEnd), endTime.Equal(timeStart), startTime.Before(timeStart), endTime.After(timeEnd), endTime, timeEnd)

			if startTime.Equal(timeEnd) || endTime.Equal(timeStart) || startTime.Before(timeStart) || endTime.After(timeEnd) {
				bResponse.HttpStatus = 400
				bResponse.StatusCode = "APP-400-005"
				bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
				bResponse.Data = nil
				log.Error().Fields(logData).Caller().Msg(bResponse.StatusDesc)
				return bResponse
			}

		}

		appointmentData := appointment.AppointmentInfo{
			DoctorId:             doctorInfo.DoctorId,
			DoctorRefId:          doctorInfo.Id,
			DoctorName:           doctorInfo.Title + " " + doctorInfo.FirstName + " " + doctorInfo.LastName,
			PatientId:            patientInfo.Id,
			PatientName:          patientInfo.Title + " " + patientInfo.FirstName + " " + patientInfo.MiddleName + " " + patientInfo.LastName,
			PatientMobileNo:      patientInfo.MobileNo,
			HN:                   patientInfo.HN,
			SlotCode:             appointmentReq.SlotCode,
			AppointmentDate:      day,
			AppointmentTimeStart: startTime,
			AppointmentTimeEnd:   endTime,
			UpdatedBy:            appointmentReq.MobileNo,
			UpdatedDate:          time.Now(),
			Status:               "A",
			Remark:               nil,
		}
		log.Debug().Msgf("appointmentData : %v", appointmentData)
		rowsAffects, err := appointment.DBSaveAppointment(appointmentData)
		if err != nil {
			// log.Err(err).Fields(logData).Caller().Msg(err.Error())
			bResponse.HttpStatus = 500
			bResponse.StatusCode = "APP-500"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = err.Error()
			return bResponse
		}
		log.Debug().Msgf("%v", rowsAffects)
		if rowsAffects == -1 {
			bResponse.HttpStatus = 500
			bResponse.StatusCode = "APP-500"
			bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
			bResponse.Data = nil
			return bResponse
		}

		bResponse.HttpStatus = 201
		bResponse.StatusCode = "APP-200"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)

	}
	return bResponse
}

func CancelAppointment(c *fiber.Ctx) base.Response {
	logData := make(map[string]interface{})
	logData["tag"] = "appointment.CancelAppointment"
	bResponse := base.Response{}
	id, err := strconv.Atoi(c.Query("id", "-1"))
	if err != nil {
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}

	if id == -1 { // ไม่ได้ใส่ query param มา
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		return bResponse
	}
	apnInfo, err := appointment.DBGetAppointmentById(id)
	if apnInfo.Id <= 0 {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-404-004"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		return bResponse
	}

	if strings.ToUpper(apnInfo.Status) == "I" {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-404-004"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		return bResponse
	}
	apnInfo.Status = "I"
	apnInfo.UpdatedDate = time.Now()
	apnInfo.UpdatedBy = apnInfo.PatientMobileNo
	log.Debug().Msgf("apnInfo : %v", apnInfo)
	rowsAffects, err := appointment.DBSaveAppointment(apnInfo)
	if err != nil {
		// log.Err(err).Fields(logData).Caller().Msg(err.Error())
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}
	log.Debug().Msgf("%v", rowsAffects)
	if rowsAffects == -1 {
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		return bResponse
	}

	bResponse.HttpStatus = 201
	bResponse.StatusCode = "APP-200"
	bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
	return bResponse
}

func GetDoctorAppointmentReport(c *fiber.Ctx) base.Response {
	logData := make(map[string]interface{})
	logData["tag"] = "appointment.GetDoctorAppointmentReport"
	bResponse := base.Response{}
	id := c.Query("id", "000")

	if id == "000" { // ไม่ได้ใส่ query param มา
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-400-001"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		return bResponse
	}

	docInfo, err := doctor.DBGetDoctorInfo(id)
	if err != nil {
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}
	if docInfo.Id <= 0 {
		bResponse.HttpStatus = 400
		bResponse.StatusCode = "APP-404-002"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = nil
		return bResponse
	}

	scheduleInfo := doctor.DBGetDoctorScheduleInfo(id, []int{0, 1, 2, 3, 4, 5, 6})
	appointmentInfos, err := appointment.DBGetAppointmentByDoctorId(docInfo.Id, []string{}) //หาทุกสถานะ
	if err != nil {
		bResponse.HttpStatus = 500
		bResponse.StatusCode = "APP-500"
		bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
		bResponse.Data = err.Error()
		return bResponse
	}
	doctorAppointmentReportListRes := DoctorAppointmentReportListRes{DoctorInfo: docInfo, Schdule: scheduleInfo}
	for _, v := range appointmentInfos {
		appointmentInfoItem := AppointmentInfoItem{HN: v.HN,
			PatientName:          v.PatientName,
			PatientMobileNo:      v.PatientMobileNo,
			AppointmentDate:      v.AppointmentDate,
			AppointmentTimeStart: v.AppointmentTimeStart,
			AppointmentTimeEnd:   v.AppointmentTimeEnd,
			Status:               v.Status,
			UpdatedBy:            v.UpdatedBy,
			UpdatedDate:          v.UpdatedDate,
			Remark:               v.Remark,
			SlotCode:             v.SlotCode}
		doctorAppointmentReportListRes.Appointment = append(doctorAppointmentReportListRes.Appointment, appointmentInfoItem)
	}
	bResponse.Data = doctorAppointmentReportListRes
	bResponse.HttpStatus = 200
	bResponse.StatusCode = "APP-200"
	bResponse.StatusDesc = viper.GetString("status." + bResponse.StatusCode)
	return bResponse
}
