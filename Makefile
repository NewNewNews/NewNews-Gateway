PROTO_DIR := internal/proto
GOOGLE_APIS_DIR := $(PROTO_DIR)/google/api

.PHONY: all clean proto

all: proto

clean:
	rm -rf $(GOOGLE_APIS_DIR)
	rm -f $(PROTO_DIR)/*.pb.go

$(GOOGLE_APIS_DIR):
	mkdir -p $(GOOGLE_APIS_DIR)
	curl -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o $(GOOGLE_APIS_DIR)/annotations.proto
	curl -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o $(GOOGLE_APIS_DIR)/http.proto

proto: $(GOOGLE_APIS_DIR)
	protoc -I . \
		-I$(PROTO_DIR) \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
		$(PROTO_DIR)/service.proto $(PROTO_DIR)/news_service.proto $(PROTO_DIR)/audio.proto $(PROTO_DIR)/summary.proto $(PROTO_DIR)/compare.proto

start:
	go run cmd/server/main.go

db:
	go run github.com/steebchen/prisma-client-go generate
	go run github.com/steebchen/prisma-client-go db push
# For Windows PowerShell
# proto-windows: $(GOOGLE_APIS_DIR)
# 	protoc -I . `
# 		-I$(PROTO_DIR) `
# 		--go_out . --go_opt paths=source_relative `
# 		--go-grpc_out . --go-grpc_opt paths=source_relative `
# 		--grpc-gateway_out . --grpc-gateway_opt paths=source_relative `
# 		$(PROTO_DIR)/service.proto $(PROTO_DIR)/news_service.proto