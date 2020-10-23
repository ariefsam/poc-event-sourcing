package repository

import (
	"encoding/json"

	"github.com/keimoon/gore"
)

type LedisEvent struct {
	Conn *gore.Conn
}

func (l *LedisEvent) Write(event map[string]interface{}) (id string, err error) {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return
	}
	temp, err := gore.NewCommand("RPUSH", "write", eventJson).Run(l.Conn)
	if err != nil {
		return
	}
	id, err = temp.String()
	if err != nil {
		return
	}

	return
}

func (l *LedisEvent) ReadAfter(afterID string) (event []map[string]interface{}, err error) {
	return
}
