package queries

import (
	"errors"
	"time"

	"github.com/YahiaJouini/careflow/internal/db"
	"github.com/YahiaJouini/careflow/internal/db/models"
)

type AppointmentRequest struct {
	DoctorID        uint      `json:"doctorId"`
	AppointmentDate time.Time `json:"appointmentDate"`
	Reason          string    `json:"reason"`
}

type AppointmentUpdateRequest struct {
	AppointmentDate time.Time `json:"appointmentDate"`
	Reason          string    `json:"reason"`
}

func CreateAppointment(patientID uint, req AppointmentRequest) (*models.Appointment, error) {
	var doctor models.Doctor

	if err := db.Db.First(&doctor, req.DoctorID).Error; err != nil {
		return nil, errors.New("Doctor not found")
	}
	if !doctor.IsAvailable {
		return nil, errors.New("Doctor is currently unavailable")
	}

	appointment := models.Appointment{
		PatientID:       patientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: req.AppointmentDate,
		Reason:          req.Reason,
		Status:          models.StatusPending,
	}

	if err := db.Db.Create(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func GetPatientAppointments(patientID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment

	err := db.Db.Preload("Doctor").Preload("Doctor.User").
		Where("patient_id = ?", patientID).
		Order("appointment_date desc").
		Find(&appointments).Error

	return appointments, err
}

func UpdateAppointment(appointmentID uint, patientID uint, req AppointmentUpdateRequest) (*models.Appointment, error) {
	var appointment models.Appointment

	if err := db.Db.Where("id = ? AND patient_id = ?", appointmentID, patientID).First(&appointment).Error; err != nil {
		return nil, errors.New("Appointment not found")
	}

	if !appointment.AppointmentDate.Equal(req.AppointmentDate) {
		if appointment.Status == models.StatusConfirmed {
			appointment.Status = models.StatusPending
		}
	}

	appointment.AppointmentDate = req.AppointmentDate
	appointment.Reason = req.Reason

	if err := db.Db.Save(&appointment).Error; err != nil {
		return nil, err
	}

	return &appointment, nil
}

func CancelAppointment(appointmentID uint, patientID uint) error {
	var appointment models.Appointment

	if err := db.Db.Where("id = ? AND patient_id = ?", appointmentID, patientID).First(&appointment).Error; err != nil {
		return errors.New("Appointment not found")
	}

	appointment.Status = models.StatusCancelled
	return db.Db.Save(&appointment).Error
}

func GetPatientByUserID(userID uint) (*models.Patient, error) {
	var patient models.Patient
	if err := db.Db.Preload("User").Where("user_id = ?", userID).First(&patient).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

type UpdatePatientBody struct {
	Height            *float64  `json:"height" validate:"omitempty,gte=0"`
	Weight            *float64  `json:"weight" validate:"omitempty,gte=0"`
	BloodType         *string   `json:"bloodType" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	ChronicConditions *[]string `json:"chronicConditions" validate:"omitempty"`
	Allergies         *[]string `json:"allergies" validate:"omitempty"`
	Medications       *[]string `json:"medications" validate:"omitempty"`
}

func UpdatePatient(userID uint, body UpdatePatientBody) (*models.Patient, error) {
	var patient models.Patient

	if err := db.Db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		return nil, errors.New("patient not found")
	}

	if body.Height != nil {
		patient.Height = *body.Height
	}
	if body.Weight != nil {
		patient.Weight = *body.Weight
	}
	if body.BloodType != nil {
		patient.BloodType = *body.BloodType
	}
	if body.ChronicConditions != nil {
		patient.ChronicConditions = *body.ChronicConditions
	}
	if body.Allergies != nil {
		patient.Allergies = *body.Allergies
	}
	if body.Medications != nil {
		patient.Medications = *body.Medications
	}

	if err := db.Db.Save(&patient).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}
