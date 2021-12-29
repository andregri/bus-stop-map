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
To run redis docker container
```
sudo docker run --name redis -d -p 6379:6379 redis
```
