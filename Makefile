include .env
export

test:
	go clean -testcache; go test -v ./...


remove:
	docker rm -f twitter-bot
	docker rmi twitter-tools_app
