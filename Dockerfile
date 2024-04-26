FROM golang:1.22.2-alpine3.18 as go-builder

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

RUN apk add --no-cache $PACKAGES

RUN make build BECH32_PREFIX=ethm

FROM alpine:3.16.1

RUN apk add curl jq bash vim 

COPY --from=go-builder /app/build/rollapp-evm /usr/local/bin/rollappd

WORKDIR /app

EXPOSE 26657 1317