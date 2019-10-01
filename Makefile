run-http-server:
	go run cmd/server/http_server.go

run-grpc-server:
	go run cmd/server/grpc_server.go

run-grpc-client:
	go run cmd/client/grpc_client.go

db-init:
	psql postgres < scripts/000.sql

db-create-table:
	 psql 'host=localhost user=calendar password=123123 dbname=calendar' < scripts/001.sql

db-set-index:
	psql 'host=localhost user=calendar password=123123 dbname=calendar' < scripts/002.sql