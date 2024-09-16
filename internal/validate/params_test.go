package validate_test

import (
	"testing"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/validate"
	"github.com/stretchr/testify/assert"
)

func TestValidateParams(t *testing.T) {
	testCases := []struct {
		name string
		in   map[string]string
	}{
		{
			name: "valid params",
			in: map[string]string{
				"bitID":       "2599da85-8a05-4c2f-bd4a-755c21cd788e",
				"tenderID":    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
				"decision":    "approved",
				"serviceType": "Construction",
				"feedback":    "10",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.ValidateParams(tc.in)

			assert.NoError(t, err)
		})
	}
}
