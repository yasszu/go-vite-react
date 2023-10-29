.PHONEY: build project and run server

.PHONY: run
run:
	@docker compose up

.PHONY: stop
stop:
	@docker compose stop

.PHONY: logs
logs:
	@docker compose logs -f

.PHONY: lint
lint:
	@docker compose run --rm node npm run lint

.PHONY: prettier
prettier:
	@docker compose run --rm node npm run prettier

.PHONY: prettier-fix
prettier-fix:
	@docker compose run --rm node npm run prettier:fix
