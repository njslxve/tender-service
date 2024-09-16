package validate

import "github.com/go-playground/validator/v10"

var (
	valid = map[string]string{
		"bidId":       "uuid",
		"tenderId":    "uuid",
		"decision":    "oneof=approved rejected",
		"serviceType": "oneof=Construction Delivery Manufacture",
		"status":      "oneof=Created Published Closed",
		"feedback":    "gt=0,lte=1000",
	}
)

func ValidateParams(params map[string]string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	for key, value := range params {
		err := validate.Var(value, valid[key])
		if err != nil {
			return err
		}
	}

	return nil
}
