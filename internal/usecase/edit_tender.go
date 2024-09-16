package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/entity"
)

func (u *Usecase) EditTender(tenderId string, username string, req dto.TenderRequest) (dto.TenderResponse, error) {
	const op = "usecase.EditTender"

	old, err := u.db.GetTenderByID(tenderId)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	var new entity.Tender = old

	if req.Name != "" {
		new.Name = req.Name
	}

	if req.Description != "" {
		new.Description = req.Description
	}

	if req.ServiceType != "" {
		new.ServiceType = req.ServiceType
	}

	new.CreatorUsername = username

	new.Version++

	new.CreatedAt = time.Now()

	t, err := u.db.EditTender(new)
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
