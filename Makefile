SERVER_PATH=./cmd/server
CLIENT_PATH=./cmd/client

build-server:
	@rm -f $(SERVER_PATH)/main && go build -o $(SERVER_PATH)/main $(SERVER_PATH)/main.go
build-client:
	@rm -f $(CLIENT_PATH)/time-client && go build -o $(CLIENT_PATH)/time-client $(CLIENT_PATH)/main.go
test:
	go test ./...