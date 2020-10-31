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

func GetEventServiceMock() (mockEventService EventServiceMock) {

	events := []map[string]interface{}{
		map[string]interface{}{
			"EventName": "UserCreated",
			"Payload": map[string]interface{}{
				"ID":    "j10",
				"Name":  "Arief Hidayatulloh",
				"Phone": "0852163772",
				"Email": "arief.hidayatulloh@gmail.com",
				"Role":  "User",
				"CreatedBy": map[string]interface{}{
					"ID":    "j1",
					"Name":  "Nam Indra",
					"Email": "nam@jojonomic.com",
					"Role":  "Admin",
				},
				"CreatedAtTimestamp": 1603235393,
			},
			"SystemUnixNano": 1603235393816625000,
		},
		map[string]interface{}{
			"EventName": "UserUpdated",
			"Payload": map[string]interface{}{
				"ID":    "j10",
				"Name":  "Arief Hidayatulloh S.Komp",
				"Phone": "085212345678",
				"Email": "arief.hidayatulloh@gmail.com",
				"Role":  "User",
				"UpdatedBy": map[string]interface{}{
					"ID":    "j1",
					"Name":  "Nam Indra",
					"Email": "nam@jojonomic.com",
					"Role":  "Admin",
				},
				"CreatedAtTimestamp": 1603235414,
			},
			"SystemUnixNano": 1603235414254786000,
		},
	}

	mockEventService.On("Write", mock.MatchedBy(func(e map[string]interface{}) bool { return true })).Return("theID", nil)

	mockEventService.On("ReadAfter", mock.MatchedBy(func(id string) bool { return true })).Return(events, nil)
	return
}

func GetSampleEvent() (data map[string]interface{}) {
	data = map[string]interface{}{
		"EventName": "UserCreated",
		"Payload": map[string]interface{}{
			"ID":    "j10",
			"Name":  "Arief Hidayatulloh",
			"Email": "arief.hidayatulloh@gmail.com",
			"Role":  "User",
			"CreatedBy": map[string]interface{}{
				"ID":    "j1",
				"Name":  "Nam Indra",
				"Email": "nam@jojonomic.com",
				"Role":  "Admin",
			},
		},
	}
	return
}
