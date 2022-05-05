package database

type DB interface {
	GetQuery() string
	ListUsers() ([]string, error)
	CreateUser(name string) (string, error)
}

func PostgresDBQuery() string {
	return "PostgresDBQuery"
}
