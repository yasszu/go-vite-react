.PHONEY: build project and run server

npm-install:
	@cd web/vite-project && npm install

npm-build:
	@cd web/vite-project && npm run build

npm-watch:
	@cd web/vite-project && npm run watch

npm-lint:
	@cd web/vite-project && npm run lint

build:
	docker compose run --rm node bash -c "make npm-install"
	docker compose run --rm node bash -c "make npm-build"

run:
	@make build
	@docker-compose up

stop:
	@docker-compose stop

lint:
	docker compose run --rm node bash -c "make npm-lint"
