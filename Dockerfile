FROM golang:1.13-alpine3.12

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod tidy
RUN go build
CMD ["/app/bot_anders14_"]
