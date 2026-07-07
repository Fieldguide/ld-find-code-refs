FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -mod=vendor -o /ld-find-code-refs ./cmd/ld-find-code-refs

FROM alpine:3.22.2

# bash for command alias scripts run from coderefs.yaml
RUN apk add --no-cache git openssh bash

COPY --from=builder /ld-find-code-refs /usr/local/bin/ld-find-code-refs
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
