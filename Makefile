build:
	@go build  -o bin/go_disign_patterns
run: build
	@./bin/go_disign_patterns
test :
	@go test -v ./...

