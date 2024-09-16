package validate

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateTender(req dto.TenderRequest) error {

	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(req)
}

func ValidateBid(req dto.BidRequest) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(req)
}
