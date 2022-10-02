
run_all:
	docker-compose up


run_app:
	docker-compose rm -s -v app
	docker-compose up -d app

run_pg:
	docker-compose up -d psql

swag:
	swag init -g cmd/main/ --parseInternal --parseDependency