include .env
export

test:
	go test -v ./internal/postgresql/*



rundb:
	docker run --name twitterdb \
		-e POSTGRES_USER=fidesy \
		-e POSTGRES_PASSWORD=secrli82kh76!hyO1 \
		-e POSTGRES_DB=twitterdb \
		-dp 5432:5432 postgres