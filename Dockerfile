# syntax=docker/dockerfile:1
FROM golang:alpine AS builder
WORKDIR /app
COPY ./ ./
RUN CGO_ENABLED=0 go build -o proxy .

FROM scratch
COPY --from=builder /app/proxy /proxy
ENTRYPOINT [ "/proxy" ]
