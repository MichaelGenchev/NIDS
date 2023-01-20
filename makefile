init-db:
	go run initdb/main.go --uri=$(uri)

run:
	go run main.go
