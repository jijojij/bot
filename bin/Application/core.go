package Application

import (
	"strings"
	"tg/bin/Application/Executor"
	"tg/bin/Core"
	Mode "tg/bin/Core/mode"
	"tg/bin/Storage"
)

var exec = &Executor.Executor{}

type SomeMessage struct {
	Message  string
	Keyboard bool
}

func Apple(id int64, input string) SomeMessage {
	data := storage.Get(id)

	switch data.PreviousState {

	case Mode.Waiting:
		if input == "add" {
			exec.NewState(id, Mode.AddOrigin)
			return SomeMessage{
				Message:  "введи ориджин",
				Keyboard: false,
			}
		}

	case Mode.AddOrigin:
		event := &core.AddWordEvent{
			NewWord: strings.TrimSpace(input),
			ChatId:  id,
		}
		exec.NewWord(event)
		exec.NewState(id, Mode.AddTranslate)
		return SomeMessage{
			Message:  "введи перевод",
			Keyboard: false,
		}

	case Mode.AddTranslate:
		event := &core.AddTranslateEvent{
			NewTranslate: strings.TrimSpace(input),
			ChatId:       id,
		}
		exec.NewWord()

	default:
		panic("unhandled default case")
	}

	return SomeMessage{
		Message:  "asd",
		Keyboard: true,
	}
}
