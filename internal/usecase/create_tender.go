package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
)

func (u *Usecase) CreateTender(req dto.TenderRequest) (dto.TenderResponse, error) {
	const op = "usecase.CreateTender"

	err := u.foundUser(req.CreatorUsername)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}

	err = u.isResponsible(req.CreatorUsername, req.OrganizationId)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, ErrNotPermissions)
	}

	new := entity.Tender{
		Name:            req.Name,
		Description:     req.Description,
		ServiceType:     req.ServiceType,
		OrganizationID:  req.OrganizationId,
		CreatorUsername: req.CreatorUsername,
	}

	fromDB, err := u.db.CreateTender(new)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	tender := dto.TenderResponse{
		ID:          fromDB.ID,
		Name:        fromDB.Name,
		Description: fromDB.Description,
		Status:      fromDB.Status,
		ServiceType: fromDB.ServiceType,
		Version:     fromDB.Version,
		CreatedAt:   fromDB.CreatedAt.Format(time.RFC3339),
	}

	return tender, nil
}
