

rundb:
	docker run --name twitterdb \
		-e POSTGRES_USER=fidesy \
		-e POSTGRES_PASSWORD=secrli82kh7 \
		-e POSTGRES_DB=twitter \
		-dp 5432:5432 postgres