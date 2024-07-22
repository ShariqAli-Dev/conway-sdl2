build:
	@go build -o ./bin/conway ./cmd/conway
run: build
	@./bin/conway