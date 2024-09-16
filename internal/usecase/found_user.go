package usecase

import "log/slog"

func (u *Usecase) foundUser(username string) error {
	const op = "usecase.FoundUser"

	err := u.db.FoundUser(username)
	if err != nil {
		u.logger.Error(op, slog.String("error", ErrUserNotFound.Error()))

		return ErrUserNotFound
	}

	return nil
}
