# syntax=docker/dockerfile:1
FROM golang:1.17 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /dist
WORKDIR /dist
COPY --from=builder /app/main .
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/favicon.ico ./favicon.ico
EXPOSE 8005
CMD ["./main"]
