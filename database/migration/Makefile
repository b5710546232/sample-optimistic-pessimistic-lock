-include .env

# define default target
STEP = 1
DATABASE_URL="postgres://postgres:password@localhost:5433/postgres?sslmode=disable"


up:
	migrate -source file://./ -database $(DATABASE_URL) up $(STEP)

down:
	migrate -source file://./ -database $(DATABASE_URL) down $(STEP)
force:
	migrate -source file://./ -database $(DATABASE_URL) force $(STEP)

create-migrate:
	 migrate create -ext sql -dir ./ -seq ${NAME}