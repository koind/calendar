# calendar
Microservice calendar on Go

## Run HTTP Server

```
CALENDAR_HOST="localhost" CALENDAR_PORT=7777 go run cmd/server/http_server.go
```

## Run gRPC Server
```
CALENDAR_HOST="0.0.0.0" CALENDAR_PORT=50051 go run cmd/server/grpc_server.go
```

## API requests

##### Create new event

```bash
curl -X "POST" "http://localhost:7777/event" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
        "uuid": 123,
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
curl -X "PUT" "http://localhost:7777/event/123" \
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
curl -X "DELETE" "http://localhost:7777/event/123"
```

Result:

```json
ok
```
---

