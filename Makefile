build:
	@mkdir -p bin
	@go build -o bin/gdfs

run: build	
	@./bin/gdfs

test:
	@go test ./... -v