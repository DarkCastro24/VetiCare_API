package validators

import (
	"VetiCare/entities/dto"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidPetID           = errors.New("pet_id es obligatorio y debe ser un UUID válido")
	ErrInvalidVetID           = errors.New("vet_id debe ser un UUID válido si se proporciona")
	ErrInvalidDateOnly        = errors.New("date_only es obligatorio y debe tener formato DD-MM-YYYY")
	ErrInvalidTimeOnly        = errors.New("time_only es obligatorio y debe tener formato HH.MM")
	ErrInvalidDateTimeInPast  = errors.New("la fecha y hora de la cita no pueden ser en el pasado")
	ErrInvalidWeight          = errors.New("weight_kg debe ser un número positivo si se proporciona")
	ErrInvalidTemperature     = errors.New("temperature debe ser un número positivo si se proporciona")
	ErrInvalidReasonLength    = errors.New("reason debe tener máximo 300 caracteres")
	ErrInvalidVaccinationLen  = errors.New("vaccination_status debe tener máximo 500 caracteres")
	ErrInvalidMedicationsLen  = errors.New("medications_prescribed debe tener máximo 300 caracteres")
	ErrInvalidAdditionalNotes = errors.New("additional_notes debe tener máximo 500 caracteres")
)

func ValidateUUIDRequired(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.New("UUID inválido: " + id)
	}
	return nil
}

func ValidateUUIDOptional(id *string) error {
	if id == nil || *id == "" {
		return nil
	}
	_, err := uuid.Parse(*id)
	if err != nil {
		return errors.New("UUID inválido: " + *id)
	}
	return nil
}

func ValidateDateOnly(dateOnly string) error {
	if len(dateOnly) != 10 {
		return ErrInvalidDateOnly
	}
	_, err := time.Parse("02-01-2006", dateOnly)
	if err != nil {
		return ErrInvalidDateOnly
	}
	return nil
}

func ValidateTimeOnly(timeOnly string) error {
	if len(timeOnly) != 5 {
		return ErrInvalidTimeOnly
	}
	_, err := time.Parse("15:04", timeOnly)
	if err != nil {
		return ErrInvalidTimeOnly
	}
	return nil
}

func ValidateDateTimeNotPast(dateOnly, timeOnly string) error {
	dateParsed, err := time.Parse("02-01-2006", dateOnly)
	if err != nil {
		return ErrInvalidDateOnly
	}
	timeParsed, err := time.Parse("15.04", timeOnly)
	if err != nil {
		return ErrInvalidTimeOnly
	}
	dateTime := time.Date(dateParsed.Year(), dateParsed.Month(), dateParsed.Day(),
		timeParsed.Hour(), timeParsed.Minute(), 0, 0, time.UTC)

	if dateTime.Before(time.Now()) {
		return ErrInvalidDateTimeInPast
	}
	return nil
}

func ValidateStatusID(statusID int) error {
	if statusID < 1 || statusID > 3 {
		return errors.New("Status_id inválido, debe ser un valor numerico")
	}
	return nil
}

func ValidatePositiveFloatOptional(value *float64, errMsg error) error {
	if value != nil && *value < 0 {
		return errMsg
	}
	return nil
}

func ValidateMaxLenOptional(s string, max int, errMsg error) error {
	if len(s) > max {
		return errMsg
	}
	return nil
}

func ValidateAppointmentDTO(app dto.AppointmentDTO) error {
	if err := ValidateUUIDRequired(app.PetID); err != nil {
		return ErrInvalidPetID
	}
	if err := ValidateUUIDOptional(app.VetID); err != nil {
		return ErrInvalidVetID
	}
	if err := ValidateDateOnly(app.Date); err != nil {
		return err
	}
	if err := ValidateTimeOnly(app.Time); err != nil {
		return err
	}
	if err := ValidateDateTimeNotPast(app.Date, app.Time); err != nil {
		return err
	}
	if err := ValidateStatusID(app.StatusID); err != nil {
		return err
	}
	if err := ValidatePositiveFloatOptional(app.WeightKg, ErrInvalidWeight); err != nil {
		return err
	}
	if err := ValidatePositiveFloatOptional(app.Temperature, ErrInvalidTemperature); err != nil {
		return err
	}
	if err := ValidateMaxLenOptional(app.Reason, 300, ErrInvalidReasonLength); err != nil {
		return err
	}
	if err := ValidateMaxLenOptional(app.VaccinationStatus, 500, ErrInvalidVaccinationLen); err != nil {
		return err
	}
	if err := ValidateMaxLenOptional(app.MedicationsPrescribed, 300, ErrInvalidMedicationsLen); err != nil {
		return err
	}
	if err := ValidateMaxLenOptional(app.AdditionalNotes, 500, ErrInvalidAdditionalNotes); err != nil {
		return err
	}
	return nil
}
