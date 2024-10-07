package words

import "bot/internal/domain/users"

// UserRepository - интерфейс для работы с пользователями в хранилище данных
type UserRepository interface {
	Create(user *users.User) error              // Сохранение нового пользователя
	FindByID(userID int64) (*users.User, error) // Поиск пользователя по ID
}
