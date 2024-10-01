package Executor

import (
	"tg/bin/Core"
	Mode "tg/bin/Core/mode"
	"tg/bin/Storage"
)

type Executor struct{}

func (_ *Executor) NewWord(event *core.AddWordEvent) {
	data := storage.Get(event.ChatId)
	data.Words[event.NewWord] = ""
	storage.Save(event.ChatId, data)
}

func (_ *Executor) NewTranslate(event *core.AddWordEvent) {
	data := storage.Get(event.ChatId)
	data.Words[event.NewWord] = ""
	storage.Save(event.ChatId, data)
}

func (_ *Executor) NewState(id int64, mode Mode.Mode) {
	data := storage.Get(id)
	data.PreviousState = mode
	storage.Save(id, data)
}
