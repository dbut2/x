FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /bin/vanity-example

FROM alpine:latest
COPY --from=builder /bin/vanity-example /vanity-example
EXPOSE 8080
ENTRYPOINT ["/vanity-example"]
