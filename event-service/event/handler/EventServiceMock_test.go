package handler_test

import "github.com/stretchr/testify/mock"

type EventServiceMock struct {
	mock.Mock
}

func (m *EventServiceMock) Write(event map[string]interface{}) (id string, err error) {
	args := m.Called(event)
	id = args.String(0)
	err = args.Error(1)
	return
}

func (m EventServiceMock) ReadAfter(afterID string) (event []map[string]interface{}, err error) {
	args := m.Called(afterID)
	event = args.Get(0).([]map[string]interface{})
	err = args.Error(1)
	return
}

func GetEventServiceMock(events []map[string]interface{}) (mockEventService EventServiceMock) {
	mockEventService.On("Write", mock.MatchedBy(func(e map[string]interface{}) bool { return true })).Return("theID", nil)

	mockEventService.On("ReadAfter", mock.MatchedBy(func(id string) bool { return true })).Return(events, nil)
	return
}
