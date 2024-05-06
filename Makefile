include .env
STEP = 3

sql-gen:
	sqlboiler psql -c sqlboiler.toml

migrate-up:
	migrate -path database/migration -database $(DATABASE_URL) up $(STEP)

migrate-down:
	migrate -path database/migration -database $(DATABASE_URL) down $(STEP)