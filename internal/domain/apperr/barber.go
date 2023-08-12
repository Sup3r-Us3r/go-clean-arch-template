package apperr

var (
	ErrBarberFieldNameIsRequired     = NewBadRequestError("field name is required")
	ErrBarberFieldEmailIsRequired    = NewBadRequestError("field email is required")
	ErrBarberEmailIsInvalid          = NewBadRequestError("email is invalid")
	ErrBarberFieldPhoneIsRequired    = NewBadRequestError("field phone is required")
	ErrBarberFieldPasswordIsRequired = NewBadRequestError("field password is required")
	ErrBarberPasswordIsInvalid       = NewBadRequestError("password is invalid, must be at least 8 digits, 1 uppercase letter, 1 lowercase letter and 1 special character")
	ErrBarberUnableToAuthenticate    = NewBadRequestError("unable to authenticate, check your credentials and try again")
	ErrBarberNotFound                = NewNotFoundError("barber not found")
	ErrBarberAlreadyExists           = NewBadRequestError("barber already exists")
)
