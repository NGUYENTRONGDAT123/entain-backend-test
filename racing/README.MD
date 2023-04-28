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
