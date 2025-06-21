package validators

import (
	"PetVet/entities/dto"
	"errors"
)

var (
	ErrInvalidFullNameUser = errors.New("nombre completo inválido: solo letras y espacios permitidos")
	ErrInvalidDUIUser      = errors.New("DUI inválido: formato esperado ########-#")
	ErrInvalidPhoneUser    = errors.New("teléfono inválido: formato esperado ####-####")
	ErrInvalidEmailUser    = errors.New("correo inválido")
)

func ValidateUserDTO(user dto.UserDTO) error {
	if err := ValidateFullName(user.FullName); err != nil {
		return ErrInvalidFullNameUser
	}
	if err := ValidateDUI(user.DUI); err != nil {
		return ErrInvalidDUIUser
	}
	if err := ValidatePhone(user.Phone); err != nil {
		return ErrInvalidPhoneUser
	}
	if err := ValidateEmail(user.Email); err != nil {
		return ErrInvalidEmailUser
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}
	return nil
}
