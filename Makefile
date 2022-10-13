.PHONEY: build project and run server

install:
	@cd vite-project && npm install

build:
	@cd vite-project && npm run build

run :
	@go run .
