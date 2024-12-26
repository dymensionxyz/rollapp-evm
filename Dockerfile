# Use Ubuntu as the base image
FROM ubuntu:latest as go-builder

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    wget make git build-essential \
    && rm -rf /var/lib/apt/lists/*

# Download and install Go 1.21
RUN wget https://golang.org/dl/go1.22.4.linux-amd64.tar.gz && \
    tar -xvf go1.22.4.linux-amd64.tar.gz && \
    mv go /usr/local && \
    rm go1.22.4.linux-amd64.tar.gz

# Set Go environment variables
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Set the working directory
WORKDIR /app

# Download go dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Cosmwasm - Download correct libwasmvm version
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

RUN go install -v github.com/bcdevtools/devd/v2/cmd/devd@latest

# Copy the remaining files
COPY . .

RUN make build BECH32_PREFIX=ethm

FROM ubuntu:latest

RUN apt-get update -y
RUN apt-get install -y curl

COPY --from=go-builder /go/bin/devd /usr/local/bin/devd
COPY --from=go-builder /app/build/rollapp-evm /usr/local/bin/rollappd
COPY --from=go-builder /lib/libwasmvm.x86_64.so /lib/libwasmvm.x86_64.so
COPY --from=go-builder /lib/libwasmvm.aarch64.so /lib/libwasmvm.aarch64.so

WORKDIR /app

EXPOSE 26657 1317