all: lint

.PHONY: lint
lint: bin/golangci-lint
	./bin/golangci-lint run -v

bin:
	@mkdir -p ./bin

bin/golangci-lint: | bin
	@echo "Installing golangci-lint..."
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/v2.15.0/install.sh | sh -s v2.15.0
