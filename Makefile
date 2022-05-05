run:
	go run main.go

mock:
	mockgen -source=database/db.go -destination=database/mocks/db_mock.go \
-package=mocks
