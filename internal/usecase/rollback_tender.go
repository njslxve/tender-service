package usecase

import (
	"fmt"
	"strconv"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
)

func (u *Usecase) RollbackTender(tenderId string, version string, username string) (dto.TenderResponse, error) {
	const op = "usecase.RollbackTender"

	ver, err := strconv.Atoi(version)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	old, err := u.db.GetTenderByVersion(tenderId, int32(ver))
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	var new entity.Tender = old

	v, err := u.db.GetTenderLastVersion(tenderId)
	if err != nil {
		return dto.TenderResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	new.Version = v + 1
	new.ID = tenderId

	new.CreatedAt = time.Now()

	fromDB, err := u.db.EditTender(new)
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
