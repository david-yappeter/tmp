commit:
	./commit.sh

# e.g. make migrate flag="--rollback --steps=17"
# e.g. make migrate flag="--force {version}"
# e.g. make migrate flag="--force 202018249289"
migrate:
	go run -tags tools . migrate $(flag)

migrate-fresh:
	go run -tags devtools . migrate-fresh

# e.g. make seed name=table_name
seed:
	go run -tags devtools . seed $(name)

# e.g. make seed-production name=table_name
seed-production:
	go run -tags devtools . seed --production $(name)
