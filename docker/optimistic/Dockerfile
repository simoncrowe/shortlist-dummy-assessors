# syntax=docker/dockerfile:1
FROM golang:1.22 AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY cmd ./cmd
RUN go build  -o /opt/shortlist/selector ./cmd/optimistic/

FROM debian:stable-slim
COPY --from=builder /opt/shortlist/selector /opt/shortlist/selector
ENTRYPOINT ["/opt/shortlist/selector"]
