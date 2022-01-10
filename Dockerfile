# syntax = docker/dockerfile:experimental
# Build Container
FROM golang:1.17.3 as builder

ENV GO111MODULE on
ENV GOPRIVATE=github.com/latonaio
WORKDIR /go/src/github.com/latonaio

COPY . .
RUN go mod download
RUN go build -o sql-update-kube ./

# Runtime Container
FROM alpine:3.14
RUN apk add --no-cache libc6-compat
ENV SERVICE=sql-update-kube \
    APP_DIR="${AION_HOME}/${POSITION}/${SERVICE}"

WORKDIR ${AION_HOME}

COPY --from=builder /go/src/github.com/latonaio/sql-update-kube .

CMD ["./sql-update-kube"]
