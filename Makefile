CONTAINER=grud_api

install:
	go get -v

setup-test:
	go get gopkg.in/check.v1

test: setup-test install
	go test ./api

run: install
	go run main.go

docker-test:
	docker exec -it ${CONTAINER} make test

docker-build:
	docker-compose up -d

docker-destroy:
	docker-compose down

docker-restart: docker-destroy
	docker-compose up -d

docker-logs:
	docker-compose logs -f api
