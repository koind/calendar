run-http-server:
	go run main.go http_server

run-grpc-server:
	go run main.go grpc_server

run-grpc-client:
	go run main.go grpc_client

db-init:
	psql postgres < scripts/000.sql

db-create-table:
	 psql 'host=localhost user=calendar password=123123 dbname=calendar' < scripts/001.sql

db-set-index:
	psql 'host=localhost user=calendar password=123123 dbname=calendar' < scripts/002.sql