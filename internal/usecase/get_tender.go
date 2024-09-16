package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
)

func (u *Usecase) GetTenders(serviceType string, limit string, offset string) ([]dto.TenderResponse, error) {
	const op = "usecase.GetTenders"

	tenders := make([]dto.TenderResponse, 0)

	fromDB, err := u.db.GetTenders(serviceType, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for _, t := range fromDB {
		tenders = append(tenders, dto.TenderResponse{
			ID:          t.ID,
			Name:        t.Name,
			Description: t.Description,
			Status:      t.Status,
			ServiceType: t.ServiceType,
			Version:     t.Version,
			CreatedAt:   t.CreatedAt.Format(time.RFC3339),
		})
	}

	if len(tenders) > 50 {
		return tenders[:50], nil
	} else {
		return tenders, nil
	}
}
