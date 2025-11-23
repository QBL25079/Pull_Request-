FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/pr-reviewer-service ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/pr-reviewer-service /usr/local/bin/

COPY wait_for_db.sh /usr/local/bin/

COPY migrations /app/migrations

ENV DB_SOURCE postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

ENTRYPOINT ["/bin/sh", "-c"]
CMD ["/usr/local/bin/wait_for_db.sh && migrate -path /app/migrations -database $DB_SOURCE up && pr-reviewer-service"]