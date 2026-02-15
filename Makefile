.PHONY: run

run:
	@if [ ! -f .env ]; then echo ".env file not found. run: cp .env.example .env"; exit 1; fi
	@set -a; . ./.env; set +a; go run .
