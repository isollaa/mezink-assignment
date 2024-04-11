MIGRATION_STEP=1

include .env

docker_start:
	docker-compose up --build

docker_stop:
	docker-compose down

migrate_create:
	@read -p "migration name (do not use space): " NAME \
  	&& migrate create -ext sql -dir ./migration $${NAME}
