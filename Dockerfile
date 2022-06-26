FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN ["go", "build"]

FROM golang:1.17-alpine

COPY --from=builder /app/test-deploy app

COPY ./.env ./.env

EXPOSE ${PORT}

CMD ["./app"]
