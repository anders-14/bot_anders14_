build:
	go build

fmt:
	goimports -w -l .

lint:
	golint ./...

image:
	docker image rm twitch-bot && \
	docker build -t twitch-bot .

run:
	docker run --rm twitch-bot
