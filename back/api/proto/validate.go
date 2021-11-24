package proto

import validation "github.com/go-ozzo/ozzo-validation"

func (a GetAimsRequest) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Period, validation.Required, validation.Length(6, 6)),
		validation.Field(&a.UserId, validation.Required),
	)
}

func (a *AimModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Period, validation.Required, validation.Length(6, 6)),
	)
}
