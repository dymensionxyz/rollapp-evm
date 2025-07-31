# Use Ubuntu as the base image
FROM ubuntu:latest as go-builder

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    wget make git build-essential \
    && rm -rf /var/lib/apt/lists/*

# Download and install Go 1.23.6
# Download and install Go
RUN ARCH=$(dpkg --print-architecture) && \
    case ${ARCH} in \
        amd64) GOARCH=amd64 ;; \
        arm64) GOARCH=arm64 ;; \
        *) echo "Unsupported architecture: ${ARCH}" && exit 1 ;; \
    esac && \
    wget https://golang.org/dl/go1.23.6.linux-${GOARCH}.tar.gz && \
    tar -xvf go1.23.6.linux-${GOARCH}.tar.gz && \
    mv go /usr/local && \
    rm go1.23.6.linux-${GOARCH}.tar.gz

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

# Cosmwasm - Download correct libwasmvm version with better architecture detection
RUN ARCH=$(uname -m) && \
    case ${ARCH} in \
        x86_64) WASMVM_ARCH=x86_64 ;; \
        aarch64|arm64) WASMVM_ARCH=aarch64 ;; \
        *) echo "Unsupported architecture: ${ARCH}" && exit 1 ;; \
    esac && \
    WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.${WASMVM_ARCH}.a \
    -O /lib/libwasmvm_muslc.a && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm.${WASMVM_ARCH}.so \
    -O /lib/libwasmvm.${WASMVM_ARCH}.so

RUN go install -v github.com/bcdevtools/devd/v2/cmd/devd@latest

# Copy the remaining files
COPY . .

RUN make build BECH32_PREFIX=ethm

FROM ubuntu:latest

RUN apt-get update -y
RUN apt-get install -y curl

COPY --from=go-builder /go/bin/devd /usr/local/bin/devd
COPY --from=go-builder /app/build/rollapp-evm /usr/local/bin/rollappd
COPY --from=go-builder /lib/libwasmvm*.so /lib/

WORKDIR /app

EXPOSE 26657 1317