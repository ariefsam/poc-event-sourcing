package repository

import (
	"encoding/json"
	"fmt"

	"github.com/keimoon/gore"
)

type RedisEvent struct {
	Conn *gore.Conn
}

func (l *RedisEvent) Write(event map[string]interface{}) (id string, err error) {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return
	}
	temp, err := gore.NewCommand("RPUSH", "write", eventJson).Run(l.Conn)
	if err != nil {
		return
	}
	idInt, err := temp.Int()
	if err != nil {
		return
	}

	id = fmt.Sprint(idInt)

	return
}

func (l *RedisEvent) ReadAfter(afterID string) (event []map[string]interface{}, err error) {
	return
}
