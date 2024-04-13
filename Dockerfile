FROM golang:1.22-alpine as build-base

WORKDIR /app

COPY go.mod .
COPY config.yml .

RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o ./out/go-sample


FROM alpine:3.19.1
COPY --from=build-base /app/out/go-sample /app/go-sample
COPY --from=build-base /app/config.yml /app/config.yml

WORKDIR /app
EXPOSE 1323
CMD ["/app/go-sample"]