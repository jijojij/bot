package core

type AddWordEvent struct {
	NewWord string
	ChatId  int64
}

type AddTranslateEvent struct {
	NewTranslate string
	ChatId       int64
}

type ResetToWaitingEvent struct{}
