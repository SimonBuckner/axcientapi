# Client for the Axcient x360Recover API

A library for accessing the Axcient x260Recover API. The library currently
supports all of the GET method endpoints. This is very 
much a work in progress and should be used with caution.

## How to Use

To install use.

```
go get -u github.com/simonbuckner/axcientapi
```

For an example of how to build a client with this library, use this project.

- [/cmd/example/main.go](https://github.com/simonbuckner/axcientapi/cmd/example/main.go)

## Outstanding Tasks

These are things I'd like to implement, time permitting. Some will come out of 
necessity as I query additional APIs.

- Write some unit and integration tests
- Write some documentation beyond the example code
- Cover the POST method endpoints
