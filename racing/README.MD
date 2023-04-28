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
