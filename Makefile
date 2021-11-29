app_build:
	@docker pull golang
	@docker pull mysql
	@docker-compose -f docker-compose.yml build --no-cache

app_start:
	@docker-compose -f docker-compose.yml up --build

app_stop:
	@docker compose -f docker-compose.yml stop

app_remove:
	@docker-compose -f docker-compose.yml down --remove-orphans

test_run:
	@echo "running test.."
	cd ./cmd/bookstore/test &\
	go test -v ./...


