package storage

import (
	Mode "tg/bin/Core/mode"
)

type Data struct {
	Words         map[string]string
	PreviousState Mode.Mode
}

var db = make(map[int64]Data)

func Save(id int64, data Data) {
	db[id] = data
}

func Get(id int64) Data {
	v, ok := db[id]
	if !ok {
		return Data{
			Words:         make(map[string]string),
			PreviousState: Mode.Waiting,
		}
	}
	return v
}
