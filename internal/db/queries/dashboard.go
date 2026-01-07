package queries

import (
	"time"

	"github.com/YahiaJouini/careflow/internal/db"
	"github.com/YahiaJouini/careflow/internal/db/models"
)

type AdminDashboardStats struct {
	TotalUsers       int64 `json:"totalUsers"`
	TotalDoctors     int64 `json:"totalDoctors"`
	VerifiedDoctors  int64 `json:"verifiedDoctors"`
	TotalSpecialties int64 `json:"totalSpecialties"`

	UsersByRole          map[string]int64 `json:"usersByRole"`          
	AppointmentsByStatus map[string]int64 `json:"appointmentsByStatus"` 
}

type DoctorDashboardStats struct {
	TotalRevenue         float64 `json:"totalRevenue"`
	PendingRequests      int64   `json:"pendingRequests"`
	UpcomingAppointments int64   `json:"upcomingAppointments"`
	TotalPatients        int64   `json:"totalPatients"`
	CompletedVisits      int64   `json:"completedVisits"`

	AppointmentsByStatus    map[string]int64 `json:"appointmentsByStatus"`   
	AppointmentsLast7Days   map[string]int64 `json:"appointmentsLast7Days"` 
}

type PatientDashboardStats struct {
	TotalAppointments    int64 `json:"totalAppointments"`
	UpcomingAppointments int64 `json:"upcomingAppointments"`
	CompletedAppointments int64 `json:"completedAppointments"`

	AppointmentsByStatus      map[string]int64 `json:"appointmentsByStatus"`     
	AppointmentsLast6Months   map[string]int64 `json:"appointmentsLast6Months"`  
}

func GetAdminStats() (*AdminDashboardStats, error) {
	stats := &AdminDashboardStats{
		UsersByRole:          make(map[string]int64),
		AppointmentsByStatus: make(map[string]int64),
	}

	if err := db.Db.Model(&models.User{}).Count(&stats.TotalUsers).Error; err != nil {
		return nil, err
	}
	if err := db.Db.Model(&models.Doctor{}).Count(&stats.TotalDoctors).Error; err != nil {
		return nil, err
	}
	if err := db.Db.Model(&models.Doctor{}).Where("is_verified = ?", true).Count(&stats.VerifiedDoctors).Error; err != nil {
		return nil, err
	}
	if err := db.Db.Model(&models.Specialty{}).Count(&stats.TotalSpecialties).Error; err != nil {
		return nil, err
	}

	rows, err := db.Db.Model(&models.User{}).Select("role, count(*)").Group("role").Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var role string
			var count int64
			rows.Scan(&role, &count)
			stats.UsersByRole[role] = count
		}
	}

	rows, err = db.Db.Model(&models.Appointment{}).Select("status, count(*)").Group("status").Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var status string
			var count int64
			rows.Scan(&status, &count)
			stats.AppointmentsByStatus[status] = count
		}
	}

	return stats, nil
}

func GetDoctorStats(userID uint) (*DoctorDashboardStats, error) {
	stats := &DoctorDashboardStats{
		AppointmentsByStatus:  make(map[string]int64),
		AppointmentsLast7Days: make(map[string]int64),
	}
	var doctor models.Doctor

	if err := db.Db.Where("user_id = ?", userID).First(&doctor).Error; err != nil {
		return nil, err
	}

	db.Db.Model(&models.Appointment{}).
		Where("doctor_id = ? AND status = ?", doctor.ID, models.StatusPending).
		Count(&stats.PendingRequests)

	db.Db.Model(&models.Appointment{}).
		Where("doctor_id = ? AND status = ? AND appointment_date > ?", doctor.ID, models.StatusConfirmed, time.Now()).
		Count(&stats.UpcomingAppointments)

	db.Db.Model(&models.Appointment{}).
		Where("doctor_id = ? AND status = ?", doctor.ID, models.StatusCompleted).
		Count(&stats.CompletedVisits)

	stats.TotalRevenue = float64(stats.CompletedVisits) * doctor.ConsultationFee

	db.Db.Model(&models.Appointment{}).
		Where("doctor_id = ?", doctor.ID).
		Distinct("patient_id").
		Count(&stats.TotalPatients)

	// Appointments by Status
	rows, err := db.Db.Model(&models.Appointment{}).
		Where("doctor_id = ?", doctor.ID).
		Select("status, count(*)").
		Group("status").
		Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var status string
			var count int64
			rows.Scan(&status, &count)
			stats.AppointmentsByStatus[status] = count
		}
	}

	rows, err = db.Db.Model(&models.Appointment{}).
		Where("doctor_id = ? AND appointment_date >= ?", doctor.ID, time.Now().AddDate(0, 0, -7)).
		Select("to_char(appointment_date, 'YYYY-MM-DD') as date, count(*)").
		Group("date").
		Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var date string
			var count int64
			rows.Scan(&date, &count)
			stats.AppointmentsLast7Days[date] = count
		}
	}

	return stats, nil
}

func GetPatientStats(userID uint) (*PatientDashboardStats, error) {
	stats := &PatientDashboardStats{
		AppointmentsByStatus:    make(map[string]int64),
		AppointmentsLast6Months: make(map[string]int64),
	}

	var patient models.Patient
	if err := db.Db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		return nil, err
	}

	db.Db.Model(&models.Appointment{}).Where("patient_id = ?", patient.UserID).Count(&stats.TotalAppointments)
	db.Db.Model(&models.Appointment{}).Where("patient_id = ? AND status = ? AND appointment_date > ?", patient.UserID, models.StatusConfirmed, time.Now()).Count(&stats.UpcomingAppointments)
	db.Db.Model(&models.Appointment{}).Where("patient_id = ? AND status = ?", patient.UserID, models.StatusCompleted).Count(&stats.CompletedAppointments)

	rows, err := db.Db.Model(&models.Appointment{}).
		Where("patient_id = ?", patient.UserID).
		Select("status, count(*)").
		Group("status").
		Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var status string
			var count int64
			rows.Scan(&status, &count)
			stats.AppointmentsByStatus[status] = count
		}
	}

	rows, err = db.Db.Model(&models.Appointment{}).
		Where("patient_id = ? AND appointment_date >= ?", patient.UserID, time.Now().AddDate(0, -6, 0)).
		Select("to_char(appointment_date, 'YYYY-MM') as month, count(*)").
		Group("month").
		Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var month string
			var count int64
			rows.Scan(&month, &count)
			stats.AppointmentsLast6Months[month] = count
		}
	}

	return stats, nil
}
