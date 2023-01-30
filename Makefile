VERSION=0.1.0

all: lint vet test build

build: mtool

mtool:
	@cd cmd/$@ && go build -o ../../bin/$@

test:
	@go test ./...

vet:
	@go vet ./...

lint:
	@revive ./...

clean:
	@rm -rf bin

count:
	@gocloc .
