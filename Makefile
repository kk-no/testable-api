# env
export DSN=root:root@tcp(localhost:3306)/testable?parseTime=true

# Go

.PHONY: run
run:
	go run cmd/testable/main.go

.PHONY: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: mod
mod:
	go mod download

.PHONY: db/up
db/up:
	cd build; \
	docker-compose up -d

.PHONY: db/down
db/down:
	cd build; \
	docker-compose down