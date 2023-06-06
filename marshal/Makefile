.PHONY: *

run :
	@go run main.go

bench :
	@go test -run="XXX" -bench=. -benchmem ./...

gen :
	@protoc --go_out=. --go_opt=paths=source_relative dtos/types.proto
