package validators

import (
	"PetVet/entities/dto"
	"errors"
)

var (
	ErrInvalidFullNameAdmin  = errors.New("Nombre completo inválido: solo letras y espacios permitidos")
	ErrInvalidUsernameAdmin  = errors.New("Nombre de usuario inválido: debe tener entre 3 y 50 caracteres")
	ErrInvalidDUIAdmin       = errors.New("DUI inválido: formato esperado ########-#")
	ErrInvalidPhoneAdmin     = errors.New("Teléfono inválido: formato esperado ####-####")
	ErrInvalidEmailAdmin     = errors.New("Correo inválido")
	ErrInvalidPasswordAdmin  = errors.New("Contraseña inválida: mínimo 6 caracteres si se proporciona")
	ErrInvalidAdminTypeAdmin = errors.New("Tipo de administrador inválido")
)

func ValidateUsername(username string) error {
	if len(username) < 3 || len(username) > 50 {
		return ErrInvalidUsernameAdmin
	}
	return nil
}

func ValidateAdminTypeID(adminTypeID int) error {
	if adminTypeID <= 0 {
		return ErrInvalidAdminTypeAdmin
	}
	return nil
}

func ValidateAdminRegisterDTO(admin dto.AdminRegisterDTO) error {
	if err := ValidateFullName(admin.FullName); err != nil {
		return ErrInvalidFullNameAdmin
	}
	if err := ValidateUsername(admin.Username); err != nil {
		return err
	}
	if err := ValidateDUI(admin.DUI); err != nil {
		return ErrInvalidDUIAdmin
	}
	if err := ValidatePhone(admin.Phone); err != nil {
		return ErrInvalidPhoneAdmin
	}
	if err := ValidateEmail(admin.Email); err != nil {
		return ErrInvalidEmailAdmin
	}
	if err := ValidatePassword(admin.Password); err != nil {
		return ErrInvalidPasswordAdmin
	}
	if err := ValidateAdminTypeID(admin.AdminTypeID); err != nil {
		return ErrInvalidAdminTypeAdmin
	}
	return nil
}
