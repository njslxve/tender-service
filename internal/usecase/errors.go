package usecase

import "errors"

var (
	ErrNotPermissions = errors.New("недостаточно прав для выполнения действия")
	ErrUserNotFound   = errors.New("пользователь не существует или некорректен")
	ErrTenderNotFound = errors.New("тендер не найден")
	ErrBidNotFound    = errors.New("предложение не найдено")
)
