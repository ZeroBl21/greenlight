# Include variables form the .envrc file
include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_/]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: confirm
confirm: 
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: run/api
run/api: ## Run the cmd/api application
	go run ./cmd/api -db-dsn=${GREENLIGHT_DB_DSN}

.PHONY: db/psql
db/psql: ## connect to the database using psql
	psql ${GREENLIGHT_DB_DSN}

.PHONY: db/migrations/new
db/migrations/new: ## create a new database migration
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/uo
db/migration/up: confirm ## apply all up database migrations
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

.PHONY: audit
audit: vendor ## tidy dependencies and format, vet and test all code
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

.PHONY: vendor
vendor: ## tidy and vendor dependencies
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor
