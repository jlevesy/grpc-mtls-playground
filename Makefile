
.PHONY: generate
generate: clean_generate
	mkdir -p dist/
	go run ./gen/main.go

.PHONY: run_server
run_server:
	go run ./server/main.go

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: clean_generate
clean_generate:
	rm -rf dist/
