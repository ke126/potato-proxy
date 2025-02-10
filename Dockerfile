# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY ./ ./
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -o proxy .

FROM scratch
COPY --from=builder /app/proxy /proxy
ENTRYPOINT [ "/proxy" ]
