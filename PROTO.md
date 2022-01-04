
## Proto Commands

1. Set path
```
export PATH="$PATH:$(go env GOPATH)/bin"
````

2. Make proto file
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
```
