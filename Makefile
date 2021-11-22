build:
	@docker pull golang
	@docker pull mysql
	@docker-compose build

start:
	@docker-compose -f docker-compose.yml build
	@docker-compose -f docker-compose.yml up

stop:
	@docker-compose -f docker-compose.yml stop

remove:
	@docker-compose -f docker-compose.yml down --remove-orphans

lint:
	@echo "lint"

test build:
	@echo "building test.."
	docker-compose -f docker-compose.test.yml build --no-cache

test run:
	@echo "running test.."
	docker-compose -f docker-compose.test.yml up

test stop:
	@echo "stopping test.."
	docker-compose -f docker-compose.test.yml stop

test down:
	@echo "stopping test.."
	docker-compose -f docker-compose.test.yml down --remove-orphans


