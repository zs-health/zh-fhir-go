# Common targets for building, testing, and running the project

.PHONY: all build test fmt vet docs run docker

all: build

clean:
	rm -rf bin/ dist/ site/ *.exe *.test
	rm -f zh-fhir

build:
	go build -v ./...

test:
	go test ./... -v

fmt:
	go fmt ./...

vet:
	go vet ./...

# run the CLI server (requires .env or flags)
run:
	go run ./cmd/zh-fhir --server --port $${PORT:-8080}

# documentation targets (requires node modules)
docs-install:
	npm install

docs-dev:
	npm run docs:dev

docs-build:
	npm run docs:build

# Docker helpers
docker-build:
	docker build -t zh-fhir .

docker-run:
	docker run -p 8080:8080 zh-fhir --server --port 8080
