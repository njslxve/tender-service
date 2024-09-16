package handler

const (
	ErrBadRequest     = "Неверный формат запроса или его параметры"
	ErrInternal       = "Сервер временно недоступен"
	ErrUserNotFound   = "Пользователь не существует или некорректен"
	ErrTenderNotFound = "Тендер не найден"
	ErrBidNotFound    = "Предложение не найдено"
	ErrNotPermissions = "Недостаточно прав для выполнения действия"
)
