COVERFILE := coverage.txt

lint:
	gofumpt -w ./..
	golangci-lint run --fix

start_api:
	go run ./cmd/api/.

clean:
	docker-compose down

build_api: clean
	docker-compose build

start: build_api
	docker-compose up

prune_volumes:
	docker volume prune --force

test: 
	go test -v -coverprofile=$(COVERFILE) ./...

go_check_deps:
	go list -u -m -json all

go_get_deps:
	go get -u ./...