FROM golang:alpine3.20 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -o /app .

# Final Stage
FROM alpine:3.20

COPY --from=builder /app /app

CMD ["/app"]