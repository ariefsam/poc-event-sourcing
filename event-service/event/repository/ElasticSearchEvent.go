package repository

type ElasticSearchEvent struct{}

func (m *ElasticSearchEvent) Write(event map[string]interface{}) (id string, err error) {
	return
}
func (m *ElasticSearchEvent) ReadAfter(afterID string) (event []map[string]interface{}, err error) {
	return
}
