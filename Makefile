.PHONEY: build project and run server

npm-install:
	@cd vite-project && npm install

npm-build:
	@cd vite-project && npm run build

npm-watch:
	@cd vite-project && npm run watch

run:
	@docker-compose up -d

stop:
	@docker-compose stop
