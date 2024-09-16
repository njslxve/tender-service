package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/njslxve/tender-service/internal/dto"
)

func ValidateTender(req dto.TenderRequest) error {

	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(req)
}

func ValidateBid(req dto.BidRequest) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(req)
}
