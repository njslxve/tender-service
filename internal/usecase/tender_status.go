package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
)

func (u *Usecase) TenderStatus(tenderId string, username string) (string, error) {
	const op = "usecase.GetTenderStatus"

	status, err := u.db.TenderStatus(tenderId, username)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return status, nil
}

func (u *Usecase) UpdateTenderStatus(tenderId string, username string, status string) (dto.TenderResponse, error) {
	const op = "usecase.UpdateTenderStatus"

	t, err := u.db.UpdateTenderStatus(tenderId, username, status)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	tender := dto.TenderResponse{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		ServiceType: t.ServiceType,
		Version:     t.Version,
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
	}

	return tender, nil
}
