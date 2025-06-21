package services

import (
	"PetVet/entities"
	"PetVet/repositories"
	"time"
)

type AppointmentService struct {
	Repo repositories.AppointmentRepository
}

func NewAppointmentService(repo repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{Repo: repo}
}

func (s *AppointmentService) CreateAppointment(app *entities.Appointment) error {
	return s.Repo.Create(app)
}

func (s *AppointmentService) GetAppointmentByID(id string) (*entities.Appointment, error) {
	return s.Repo.GetByID(id)
}

func (s *AppointmentService) GetAllAppointments() ([]entities.Appointment, error) {
	return s.Repo.GetAll()
}

func (s *AppointmentService) UpdateAppointment(id string, fields map[string]interface{}) error {
	return s.Repo.Update(id, fields)
}

func (s *AppointmentService) GetAppointmentsByStatus(statusID int) ([]entities.Appointment, error) {
	return s.Repo.GetAppointmentsByStatus(statusID)
}

func (s *AppointmentService) GetAppointmentsByStatusAndDate(date time.Time) ([]entities.Appointment, error) {
	return s.Repo.GetAppointmentsByStatusAndDate(date)
}

func (s *AppointmentService) UpdateStatus(id string, statusID int) error {
	return s.Repo.UpdateStatus(id, statusID)
}

func (s *AppointmentService) GetByUserID(userID string) ([]entities.Appointment, error) {
	return s.Repo.GetByUserID(userID)
}

func (s *AppointmentService) GetMedicalHistoryByPetID(petID string) ([]entities.Appointment, error) {
	return s.Repo.GetMedicalHistoryByPetID(petID)
}

func (s *AppointmentService) DeleteAppointment(id string) (string, error) {
	newStatus, err := s.Repo.Delete(id)
	if err != nil {
		return "", err
	}
	if newStatus == 1 {
		return "Cita reactivada correctamente", nil
	}
	return "Cita cancelada correctamente", nil
}
