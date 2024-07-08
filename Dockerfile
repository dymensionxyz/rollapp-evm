FROM golang:1.22.4-alpine3.19 as go-builder

WORKDIR /app

RUN apk add --no-cache ca-certificates build-base git

COPY go.mod go.sum* ./

RUN ARCH=$(uname -m) && WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$ARCH.a \
    -O /lib/libwasmvm_muslc.a && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.$ARCH | cut -d ' ' -f 1) && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm.x86_64.so \
    -O /lib/libwasmvm.x86_64.so && \
     wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm.aarch64.so \
    -O /lib/libwasmvm.aarch64.so

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
