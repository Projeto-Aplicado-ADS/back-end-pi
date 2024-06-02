package errors

import (
	"errors"
	"fmt"
)

var (
	ErrCheckDataExists         = fmt.Errorf("failed to check if data already exists")
	ErrDeleteData              = fmt.Errorf("failed to delete data")
	ErrDeleteFile              = fmt.Errorf("failed to delete file from storage")
	ErrDataAlreadyExists       = fmt.Errorf("data already exists")
	ErrGetStoredData           = fmt.Errorf("failed to get stored data")
	ErrGetStorage              = fmt.Errorf("failed to get storage")
	ErrNotFoundData            = fmt.Errorf("not found data")
	ErrInvalidRequest          = fmt.Errorf("invalid request data")
	ErrInvalidRequestParameter = fmt.Errorf("invalid request parameter")
	ErrInactiveResource        = fmt.Errorf("inactive resource")
	ErrPersistData             = fmt.Errorf("failed to persist data")
	ErrPersistFile             = fmt.Errorf("failed to persist file to storage")
	ErrProviderAlreadyExists   = fmt.Errorf("provider already exists")
	ErrUserIsNotAdmin          = fmt.Errorf("user is not admin")
	ErrWithoutPermission       = fmt.Errorf("user without permission for the requested resource")
	ErrWithoutSuperPermission  = fmt.Errorf("user without super permission for the requested resource")
	ErrGetPermission           = fmt.Errorf("failed to get permission")
	ErrValidation              = fmt.Errorf("validation error")
	ErrGetUserProvider         = fmt.Errorf("error getting user from provider")
	ErrOnboardingState         = fmt.Errorf("user is not in the onboarding process")
	ErrNotEligibleOnboarding   = fmt.Errorf("user is not eligible for onboarding")
	ErrBeginTransaction        = fmt.Errorf("failed to start transaction")
	ErrCommitTransaction       = fmt.Errorf("failed to commit transaction")
	ErrInvalidInvite           = fmt.Errorf("type of unexpected invitation")

	ErrValidationWithMessage = func(message string) error {
		return fmt.Errorf("%w: %s", ErrValidation, message)
	}

	ErrGetStoredDataWithMessage = func(message string) error {
		return fmt.Errorf("%w: %s", ErrGetStoredData, message)
	}

	ErrPersistDataWithMessage = func(message string) error {
		return fmt.Errorf("%w: %s", ErrPersistData, message)
	}

	ErrDeleteDataWithMessage = func(message string) error {
		return fmt.Errorf("%w: %s", ErrDeleteData, message)
	}
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
