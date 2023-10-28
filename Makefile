.PHONEY: build project and run server

run:
	@docker compose up

lint:
	@docker compose run --rm node npm run lint
