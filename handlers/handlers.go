package handlers

import (
	"fmt"
	"github.com/Ad3bay0c/interface_testing/database"
	"log"
)

type handler struct {
	db database.DB
}
func CreateHandler(db database.DB) handler {
	return handler{
		db: db,
	}
}
func (h *handler) GetDBQuery() string {
	fmt.Println("Get DB Query")
	result := h.db.GetQuery()
	return result
}

func (h *handler) ListUsers() ([]string, error) {
	return h.db.ListUsers()
}

func (h *handler) CreateUser(val string) (string, error) {
	name := fmt.Sprintf("%v", val)

	name, err := h.db.CreateUser(name)
	log.Println(err)
	return name, err
}

