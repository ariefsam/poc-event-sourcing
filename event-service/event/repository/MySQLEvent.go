package repository

type MySQLEvent struct{}

func (m *MySQLEvent) Write(event map[string]interface{}) (id string, err error) {
	return
}
func (m *MySQLEvent) ReadAfter(afterID string) (event []map[string]interface{}, err error) {
	return
}
