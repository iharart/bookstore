app build:
	@docker pull golang
	@docker pull mysql
	@docker-compose -f docker-compose.yml build --no-cache

app start:
	@docker-compose -f docker-compose.yml up --build

app stop:
	@docker-compose -f docker-compose.yml stop

app remove:
	@docker-compose -f docker-compose.yml down --remove-orphans

lint:
	@echo "lint"

test build:
	@echo "building test.."
	docker-compose -f docker-compose.test.yml build --no-cache

test_run:
	@echo "running test.."
	docker-compose -f docker-compose.test.yml up --build
test stop:
	@echo "stopping test.."
	docker-compose -f docker-compose.test.yml stop

test down:
	@echo "stopping test.."
	docker-compose -f docker-compose.test.yml down --remove-orphans


