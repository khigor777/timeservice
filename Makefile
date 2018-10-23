SERVER_PATH=./cmd/server
run-server:
	@rm -f $(SERVER_PATH)/main && go build -o $(SERVER_PATH)/main $(SERVER_PATH)/main.go && $(SERVER_PATH)/main