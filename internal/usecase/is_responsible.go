package usecase

import (
	"log/slog"
)

func (u *Usecase) isResponsible(user string, org string) error {
	const op = "usecase.IsResponsible"

	err := u.db.IsResponsible(user, org)
	if err != nil {
		u.logger.Error(op, slog.String("error", err.Error())) // TODO: add error message

		return ErrNotPermissions
	}

	return nil
}
