package handler

type EventService interface {
	Write(event map[string]interface{}) (id string, err error)
	Read(afterID string) (id string, event []map[string]interface{}, err error)
}
