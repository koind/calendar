# api
Microservice API - provides a GRPC interface for users, implements methods for creating, modifying, searching, and deleting events.

## Run HTTP Server

```
go run main.go http_server
```

## Run gRPC Server
```
go run main.go grpc_server
```

## Run gRPC Client
```
go run main.go grpc_client
```

## Update proto file 
```
protoc --go_out=plugins=grpc:. app/transport/grpc/pb/event.proto
```

## API requests

##### Create new event

```bash
curl -X "POST" "http://localhost:7766/event" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
        "title": "test1",
        "datetime": "2019-08-18T10:00:00Z",
        "duration": 15,
        "user_id": 3,
        "time_send_notify": "2019-08-18T10:00:00Z"
      }'
```

Result:

```json
{
  "uuid": 123,
  "title": "test1",
  "datetime": "2019-08-18T10:00:00Z",
  "duration": 15,
  "user_id": 3,
  "time_send_notify": "2019-08-18T10:00:00Z"
}
```
---

##### Update event

```bash
curl -X "PUT" "http://localhost:7766/event/123" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
        "uuid": 123,
        "title": "Buy watch",
        "datetime": "2019-09-18T10:00:00Z",
        "duration": 19,
        "user_id": 3,
        "time_send_notify": "2019-09-18T10:00:00Z"
      }'
```

Result:

```json
{
  "uuid": 123,
  "title": "Buy watch",
  "datetime": "2019-09-18T10:00:00Z",
  "duration": 19,
  "user_id": 3,
  "time_send_notify": "2019-09-18T10:00:00Z"
}
```
---


##### Delete event

```bash
curl -X "DELETE" "http://localhost:7766/event/123"
```

Result:

```json
ok
```
---

