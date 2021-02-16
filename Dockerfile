FROM golang:1.15 as builder

LABEL maintainer="Dody Satria <dody.satria@gmail.com>"

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server

CMD ["/server"]

#RUN go build -mod=readonly -v -o server
#
#FROM debian:buster-slim
#RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
#    ca-certificates && \
#    rm -rf /var/lib/apt/lists/*
#
#COPY --from=builder /app/server /app/server
#
#CMD ["/app/server"]