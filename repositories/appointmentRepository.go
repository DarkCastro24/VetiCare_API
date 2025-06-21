package repositories

import (
	"PetVet/entities"
	"errors"
	"time"

	"gorm.io/gorm"
)

type appointmentRepositoryGORM struct {
	db *gorm.DB
}

func NewAppointmentRepositoryGORM(db *gorm.DB) AppointmentRepository {
	return &appointmentRepositoryGORM{db: db}
}

func (r *appointmentRepositoryGORM) Create(app *entities.Appointment) error {
	return r.db.Create(app).Error
}

func (r *appointmentRepositoryGORM) GetByID(id string) (*entities.Appointment, error) {
	var app entities.Appointment
	err := r.db.
		Preload("Pet").
		Preload("Pet.Owner").
		Preload("Pet.Species").
		Preload("Vet").
		Preload("Vet.Role").
		Where("id = ?", id).
		First(&app).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &app, err
}

func (r *appointmentRepositoryGORM) GetByUserID(userID string) ([]entities.Appointment, error) {
	var apps []entities.Appointment
	err := r.db.Joins("JOIN pets ON pets.id = appointments.pet_id").
		Where("pets.owner_id = ?", userID).
		Preload("Pet").
		Preload("Pet.Owner").
		Preload("Pet.Species").
		Preload("Vet").
		Preload("Vet.Role").
		Find(&apps).Error
	return apps, err
}

func (r *appointmentRepositoryGORM) GetAll() ([]entities.Appointment, error) {
	var apps []entities.Appointment
	err := r.db.
		Preload("Pet").
		Preload("Pet.Owner").
		Preload("Pet.Species").
		Preload("Vet").
		Preload("Vet.Role").
		Find(&apps).Error
	return apps, err
}

func (r *appointmentRepositoryGORM) GetAppointmentsByStatus(statusID int) ([]entities.Appointment, error) {
	var apps []entities.Appointment
	err := r.db.
		Where("status_id = ?", statusID).
		Preload("Pet").
		Preload("Pet.Owner").
		Preload("Pet.Species").
		Preload("Vet").
		Preload("Vet.Role").
		Find(&apps).Error
	return apps, err
}

func (r *appointmentRepositoryGORM) GetAppointmentsByStatusAndDate(date time.Time) ([]entities.Appointment, error) {
	var apps []entities.Appointment
	dateStr := date.Format("02-01-2006")
	err := r.db.
		Where("status_id = 1 or status_id = 2 AND date = ?", dateStr).
		Preload("Pet").
		Preload("Pet.Owner").
		Preload("Pet.Species").
		Preload("Vet").
		Preload("Vet.Role").
		Find(&apps).Error
	return apps, err
}

func (r *appointmentRepositoryGORM) GetMedicalHistoryByPetID(petID string) ([]entities.Appointment, error) {
	var apps []entities.Appointment
	err := r.db.Where("pet_id = ? AND status_id = ?", petID, 2).
		Preload("Pet").
		Preload("Pet.Owner").
		Preload("Pet.Species").
		Preload("Vet").
		Preload("Vet.Role").
		Find(&apps).Error
	return apps, err
}

func (r *appointmentRepositoryGORM) Update(id string, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}
	return r.db.Model(&entities.Appointment{}).Where("id = ?", id).Updates(fields).Error
}

func (r *appointmentRepositoryGORM) UpdateStatus(id string, statusID int) error {
	return r.db.Model(&entities.Appointment{}).
		Where("id = ?", id).
		Update("status_id", statusID).Error
}

func (r *appointmentRepositoryGORM) Delete(id string) (int, error) {
	var app entities.Appointment
	result := r.db.First(&app, "id = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("cita no encontrada")
	}
	newStatus := 1
	if app.StatusID == 1 {
		newStatus = 2
	}
	err := r.db.Model(&entities.Appointment{}).Where("id = ?", id).Update("status_id", newStatus).Error
	if err != nil {
		return 0, err
	}
	return newStatus, nil
}
