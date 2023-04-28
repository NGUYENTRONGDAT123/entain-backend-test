For Unit Testing

1. Go to test folder:

```
cd ./test
```

2. Run Unit test:

```
go test ./...
```

1. Adding additional Visible field in gRPC as Optional Boolean Type.

HTTP POST `v1/list-races` now returns list of races:

- Request Body Schema `application/json` fields:
  `visible` (OPTIONAL): accepts bool value, return a list of races that has visible true/false according to the `visible` value. If there is no value, return list of all races.

- Request Body Schema `application/json` example:

```
{
"field" : {
  "visible": true
  }
}
```

2. Adding additional order_by field in gRPC as Optional ENUM OrderBy field

```
Enum OrderBy {
  ASC
  DESC
}
```

HTTP POST v1/list-races now returns list of races that is ordered by advertised start time:
Request Body Schema application/json fields:
`order_by` (OPTIONAL): accepts String which is either "ASC", "DESC" which returns a list of races that is ordered by advertise start time ascending or descending depending on the value of `order_by` . If there is no value, return list of all races without order.

- Request Body Schema application/json example:

```
{
"filter" : {
  "orderBy": "ASC"
  }
}
```

3. Introduce `Status` as a new field in the Race response message in gRPC to show the status whether the race is opened or closed to check if the advertised start time is in the past

Example response of list of races in JSON data:

```
{
    "races": [
        {
            "id": "1",
            "meetingId": "5",
            "name": "North Dakota foes",
            "number": "2",
            "visible": false,
            "advertisedStartTime": "2021-03-03T01:30:57Z",
            "status": "CLOSED"
        },
        {
            "id": "2",
            "meetingId": "1",
            "name": "Connecticut griffins",
            "number": "12",
            "visible": true,
            "advertisedStartTime": "5555-03-02T19:16:58Z",
            "status": "OPEN"
        },
   ]
}
```

4. introduce a new GET method API endpoint: `v1/race/{id}` where it return values of one race get by id as path parameter:

Example GET response method `v1/race/1`: JSON data

```
{
    "id": "1",
    "meetingId": "5",
    "name": "North Dakota foes",
    "number": "2",
    "visible": false,
    "advertisedStartTime": "2021-03-03T01:30:57Z",
    "status": "CLOSED"
}
```
