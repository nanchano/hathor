.PHONY: pb sqlc migrate tidy check test up down

pb:
	protoc ./pb/hathor.proto \
		--experimental_allow_proto3_optional \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		--grpc-gateway_out . \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt paths=source_relative \
			--grpc-gateway_opt grpc_api_configuration=./pb/hathor_annotations.yml \
		--openapiv2_out . \
			--openapiv2_opt logtostderr=true \
			--openapiv2_opt generate_unbound_methods=true \

sqlc:
	sqlc generate --experimental --file ./sql/sqlc.yaml

migrate:
	migrate create -ext sql -dir ./internal/repository/migrations -seq create_events_table

tidy:
	go fmt ./...
	go mod tidy -v

check:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

test:
	go test -v -race -buildvcs ./...

up:
	docker-compose up --build

down:
	docker-compose down
