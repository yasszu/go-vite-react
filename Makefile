.PHONEY: build project and run server

npm-install:
	@cd vite-project && npm install

npm-build:
	@cd vite-project && npm run build

npm-watch:
	@cd vite-project && npm run watch

npm-lint:
	@cd vite-project && npm run lint

init:
	docker compose run --rm node bash -c "make npm-install"
	docker compose run --rm node bash -c "make npm-build"

run:
	@docker-compose up

stop:
	@docker-compose stop

lint:
	docker compose run --rm node bash -c "make npm-lint"
