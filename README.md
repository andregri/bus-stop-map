# bus-stop-map

<p align="left">
  <a href="https://github.com/andregri/bus-stop-map/actions"><img alt="GitHub Actions status" src="https://github.com/andregri/bus-stop-map/workflows/Build%20go%20app/badge.svg"></a>
</p>

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
| `/v1/arrival`  | POST | Create a new arrival time recording.Fields are required in JSON body. Return the id of the created resource in JSON. |
| `/v1/arrival/id` | GET | Read an arrival time by id. Return all the fields in JSON. |
| `/v1/arrival/id` | DELETE | Delete an arrival time by id |
| `/v1/arrival/id` | PATCH | Update the arrival time record. The time field to be updated is required in JSON. |
| `/v1/arrivals/<stop code>` | GET | Get all items whose stop_code contains "<stop_code>" |

## For developers
To run the server locally:
```
docker run --name some-postgres --rm -p 5432:5432 -e POSTGRES_USER=andrea -e POSTGRES_PASSWORD=very_strong_password -e POSTGRES_DB=app_database -d postgres

POSTGRES_HOST=127.0.0.1 POSTGRES_USER=andrea POSTGRES_PASSWORD=very_strong_password POSTGRES_DB=app_database go run cmd/main.go
```

Inspect postgres:
```
docker exec -it some-postgres bash
```

then inside docker bash:
```
psql -U andrea -d app_database
```


## Test
Create a row
```
curl -X POST http://localhost:9000/v1/arrival \
-H 'content-type: application/json' \
-d '{"bus_line":"11","stop_code":"SP123","time":"11:15"}'
```

Get a row
```
curl -X GET http://localhost:9000/v1/arrival/1
```

Delete a row:
```
curl -X DELETE http://localhost:9000/v1/arrival/1
```

Update a row:
```
curl -X PATCH http://localhost:9000/v1/arrival/2 \-H 'content-type: application/json' \
-d '{"time":"22:15"}'
```

Get rows by stop code:
```
curl -X GET http://localhost:9000/v1/arrivals/SP \-H 'content-type: application/json'
```

## Build
```
cd cmd
GOOS=linux GOARCH=amd64 go build -o bus-server
```

## Deploy stack
Validate template:
```
aws cloudformation validate-template --template-body file://$PWD/iac/ec2-postgres-template.yml
```

Deploy the stack:
```
aws cloudformation deploy \
  --template-file iac/ec2-postgres-template.yml \
  --stack-name "mystack" \
  --parameter-overrides \
  KeyName=stack-kp \
  --capabilities CAPABILITY_IAM
```

Update the stack:
```
aws cloudformation update-stack \
  --stack-name mystack \
  --template-body file://$PWD/iac/ec2-postgres-template.yml \
  --parameters \
  ParameterKey=KeyName,UsePreviousValue=true \
  --capabilities CAPABILITY_IAM
```

Delete stack:
```
aws cloudformation delete-stack --stack-name mystack
```

Troubleshoot `cfn-init`
```
sudo /opt/aws/bin/cfn-init -v --stack mystack --resource WebServer --region eu-west-1
cat /var/log/cfn-init-cmd.log
```

## ToDo
[ ] Set password to user `postgres` 
[ ] Alter role in psql?
[ ] Add `cfn-signal`?