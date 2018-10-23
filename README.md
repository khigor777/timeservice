# Time service
Time client server
Implements the Time Protocol, as described in RFC 868.

### Make commands
```
make build-server
make build-client:
make test
```

### Run  example
```
server: ./cmd/server/main -p 8080
client: ./cmd/client/time-client 127.0.0.1 8080
```


