package main

import (
	"fmt"
	"github.com/Ad3bay0c/interface_testing/database"
	"github.com/Ad3bay0c/interface_testing/handlers"
)

type name struct {
	first string
	last  string
}

func main() {
	postgresDb := database.CreatePostgresDB()
	h := handlers.CreateHandler(postgresDb)
	fmt.Println(h.GetDBQuery())
}
