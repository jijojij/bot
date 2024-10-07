package users

import "bot/internal/domain/words"

// WordRepository - интерфейс для работы со словами в хранилище данных
type WordRepository interface {
	Save(word *words.Word) error                  // Сохранение нового слова
	FindByUser(userID int64) (*words.Word, error) // Поиск слова по термину и ID пользователя
}
