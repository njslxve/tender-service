package validate_test

import (
	"testing"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/validate"
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
