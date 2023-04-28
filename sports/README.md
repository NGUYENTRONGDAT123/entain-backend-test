For Unit Testing

1. Go to test folder:

```
cd ./test
```

2. Run Unit test:

```
go test ./...
```

1. Introduce new gRPC sports service.

Introduce new POST method API endpoint `v1/list-sports-events` that returns a list of sport events that is similar to the list of race.

Request Body Schema JSON:

The body schema contains `filter` field that contains all filter value for the list of sports of events results be filtered.
At the moment, `filter` only has `ids` field. If there is an id not found, it still returns OK response but the data is not in it:

- `ids` accept an array of id of the events.

Example of JSON body schema

```
{
  "filter" : {
    "ids": [1,2]
  }
}
```

Response in JSON result:

It returns a list of sports events. Depending on the body schema, it would return a filtered result of list of sport events. Each of which contains these fields:

- `id`: The id of the sport event
- `name`: The name of the sport event
- `cityAddress`: the name of the city where the event is held
- `numOfParticipants`: The number of participants who will/ participated the sport event
- `advertisedStartTime`: The official time that the event is/was started, which is in ISO 8601 format

Example of JSON response

```
{
    "events": [
        {
            "id": "1",
            "name": "Bike Racing",
            "cityAddress": "Laviniatown",
            "numOfParticipants": "286",
            "advertisedStartTime": "2023-04-27T15:03:15Z"
        },
        {
            "id": "2",
            "name": "Human Racing",
            "cityAddress": "New Audra",
            "numOfParticipants": "640",
            "advertisedStartTime": "2023-04-26T11:56:22Z"
        }
    ]
}
```
