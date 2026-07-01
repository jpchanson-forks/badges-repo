FROM golang:1.23-alpine AS builder
WORKDIR /app

# copy only module first (important for caching)
COPY badgegen/go.mod badgegen/go.sum ./
RUN go mod download

# copy source
COPY badgegen/main.go ./
RUN go build -o /bin/badgegen .


FROM alpine:3.20

COPY --from=builder /bin/badgegen /bin/badgegen

ENTRYPOINT ["/bin/badgegen"]

