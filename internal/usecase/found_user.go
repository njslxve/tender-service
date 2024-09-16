package usecase

import "log/slog"

func (u *Usecase) foundUser(username string) error {
	const op = "usecase.FoundUser"

	err := u.db.FoundUser(username)
	if err != nil {
		u.logger.Error(op, slog.String("error", err.Error())) // TODO: add error message

		return ErrUserNotFound
	}

	return nil
}
