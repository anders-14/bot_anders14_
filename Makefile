build:
	goimports -w -l .
	go build

fmt:
	goimports -w -l .

image:
	docker image rm twitch-bot && \
	docker build -t twitch-bot .

run:
	docker run --rm twitch-bot
