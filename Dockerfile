FROM golang:1 as builder

ARG VERSION=unknown

WORKDIR /lc-api
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o ./bin/lc-api \
    -ldflags="-X main.Version=$VERSION" \
    -v ./cmd/lc-api/main.go

FROM alpine:3

ARG LC_API_DIR=/opt/lc-api
ARG LC_API_USER=lc-api

RUN apk update \
    && apk add ca-certificates tzdata \
    && rm -rf /var/cache/apk/*

RUN addgroup -S $LC_API_USER \
    && adduser -S $LC_API_USER -G $LC_API_USER \
    && mkdir -p $LC_API_DIR

COPY --from=builder /lc-api/bin/lc-api $LC_API_DIR/lc-api

RUN chown -R $LC_API_USER:$LC_API_USER $LC_API_DIR

ENV LC_API_ADDR=0.0.0.0:54010
ENV LC_API_SHUTDOWN_TIMEOUT=5

ENV LC_API_RENDERER_ADDRESS=dns:///localhost:54020
ENV LC_API_RENDERER_CONN_TIMEOUT=5
ENV LC_API_RENDERER_REQUEST_TIMEOUT=30

USER $LC_API_USER
WORKDIR $LC_API_DIR

ENTRYPOINT ["./lc-api"]
CMD []
