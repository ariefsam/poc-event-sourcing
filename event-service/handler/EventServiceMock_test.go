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
