package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
)

func (u *Usecase) GetUserTenders(username string, limit string, offset string) ([]dto.TenderResponse, error) {
	const op = "usecase.GetUserTenders"

	err := u.foundUser(username)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}

	tenders := make([]dto.TenderResponse, 0)

	fromDB, err := u.db.GetUserTenders(username, limit, offset)
	if err != nil {
		return tenders, fmt.Errorf("%s: %w", op, err)
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
