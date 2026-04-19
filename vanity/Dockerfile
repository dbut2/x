FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/vanity ./cmd/server

FROM alpine AS final

LABEL org.opencontainers.image.source="https://github.com/dbut2/x"
LABEL org.opencontainers.image.url="https://github.com/dbut2/x"
LABEL org.opencontainers.image.title="vanity"

COPY --from=builder /bin/vanity /bin/vanity

EXPOSE 8080

CMD ["/bin/vanity"]
