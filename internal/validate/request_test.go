package validate_test

import (
	"testing"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/validate"
	"github.com/stretchr/testify/assert"
)

func TestValidateTender(t *testing.T) {
	testCases := []struct {
		name string
		in   dto.TenderRequest
	}{
		{
			name: "valid tender",
			in: dto.TenderRequest{
				Name:            "test",
				Description:     "test",
				ServiceType:     "Delivery",
				OrganizationId:  "2599da85-8a05-4c2f-bd4a-755c21cd788e",
				CreatorUsername: "test",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.ValidateTender(tc.in)

			assert.NoError(t, err)
		})
	}
}

func TestValidateBid(t *testing.T) {
	testCases := []struct {
		name string
		in   dto.BidRequest
	}{
		{
			name: "valid bid",
			in: dto.BidRequest{
				Name:        "test",
				Description: "test",
				TendterID:   "2599da85-8a05-4c2f-bd4a-755c21cd788e",
				AuthorType:  "User",
				AuthorID:    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.ValidateBid(tc.in)

			assert.NoError(t, err)
		})
	}
}
