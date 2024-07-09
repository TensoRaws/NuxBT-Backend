FROM golang:1.21.12-alpine3.20 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o nuxbt .

FROM alpine:3.20 AS runner

WORKDIR /app

COPY ./conf/nuxbt.yml /app/conf/
COPY --from=builder /build/nuxbt /app

EXPOSE 8080

ENTRYPOINT ["/app/nuxbt", "server"]
