package database

type MongoDB struct {}

func CreateMongoDB() *MongoDB {
	return &MongoDB{}
}
func (m *MongoDB) GetQuery() string {
	return "mongodb"
}

func (m *MongoDB)  ListUsers() ([]string, error) {
	return nil, nil
}

func (m *MongoDB)  ListDatabases() ([]string, error) {
	return nil, nil
}