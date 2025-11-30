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
}

type DoctorDashboardStats struct {
	TotalRevenue         float64 `json:"totalRevenue"`
	PendingRequests      int64   `json:"pendingRequests"`
	UpcomingAppointments int64   `json:"upcomingAppointments"`
	TotalPatients        int64   `json:"totalPatients"`
	CompletedVisits      int64   `json:"completedVisits"`
}

func GetAdminStats() (*AdminDashboardStats, error) {
	stats := &AdminDashboardStats{}

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

	return stats, nil
}

func GetDoctorStats(userID uint) (*DoctorDashboardStats, error) {
	stats := &DoctorDashboardStats{}
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

	return stats, nil
}
