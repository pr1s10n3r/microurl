package handlers

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func init() {

}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(i any) []ErrorResponse {
	errs := make([]ErrorResponse, 0)

	if err := validate.Struct(i); err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return nil
		}

		for _, verr := range verrs {
			rerr := ErrorResponse{
				FailedField: verr.StructNamespace(),
				Tag:         verr.Tag(),
				Value:       verr.Param(),
			}

			errs = append(errs, rerr)
		}

		return errs
	}

	return nil
}
