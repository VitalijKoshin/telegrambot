build:
	docker-compose build tbot-app

run:
	docker-compose up tbot-app

test:
	go test -v ./...