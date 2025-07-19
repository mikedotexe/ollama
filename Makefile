.PHONY: vet test benchwarm

vet:
go vet ./...

test:
CGO_ENABLED=0 go test -race ./...

benchwarm:
CGO_ENABLED=0 go run ./cmd/benchwarm >warm.csv

