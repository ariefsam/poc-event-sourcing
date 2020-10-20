package handler

type EventService interface {
	Write(event map[string]interface{}) (id string, err error)
	ReadAfter(afterID string) (event []map[string]interface{}, err error)
}
