FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /reddit-api

EXPOSE 3000

CMD [ "/reddit-api"]
