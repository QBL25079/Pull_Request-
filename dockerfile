FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /pr-reviewer-service cmd/server/main.go


FROM alpine:latest
RUN apk --no-cache add tzdata

COPY --from=builder /pr-reviewer-service /pr-reviewer-service

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/pr-reviewer-service"]