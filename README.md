# bus-stop-map

## API v1

### Arrival resource

| Field | Type | Description |
| --- | --- | --- |
| ID | int | unique identifier for the arrival time |
| stopCode | string | unique code identifying the bus stop | 
| busLine | string | unique code identifying the bus line |
| time | string | time of the arrival event | 

| URL | REST Verb | Action |
| --- | --- | --- |
| `/v1/arrival`  | POST | Create a new arrival time recording.Fields are required in JSON body. Return the id of the created resource. |
| `/v1/arrival/id` | GET | Read an arrival time by id. Return all the fields in JSON. |
| `/v1/arrival/id` | DELETE | Delete an arrival time by id |
| `/v1/arrival/id` | PATCH | Update the arrival time record. The field to be updated is required in JSON. |

## For developers
To run the server locally:
```
docker run --name some-postgres --rm -p 5432:5432 -e POSTGRES_USER=andrea -e POSTGRES_PASSWORD=very_strong_password -e POSTGRES_DB=app_database -d postgres

POSTGRES_HOST=127.0.0.1 POSTGRES_USER=andrea POSTGRES_PASSWORD=very_strong_password POSTGRES_DB=app_database go run cmd/main.go
```
