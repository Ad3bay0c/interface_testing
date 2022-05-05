package database

type PostgresDB struct {}

func CreatePostgresDB() *PostgresDB {
	return &PostgresDB{}
}
func (m *PostgresDB) GetQuery() string {
	return "postgresdb"
}

func (m *PostgresDB)  ListUsers() ([]string, error) {
	return nil, nil
}

func (m *PostgresDB)  ListDatabases() ([]string, error) {
	return nil, nil
}