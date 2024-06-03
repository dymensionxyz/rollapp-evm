#!/usr/bin/make -f

PROJECT_NAME=rollappd
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true

ifndef BECH32_PREFIX
    $(error BECH32_PREFIX is not set)
endif

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --tags)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

export GO111MODULE = on

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=dymension-rdk \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=rollapp-evm \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	      -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION) \
		  -X github.com/dymensionxyz/rollapp-evm/app.AccountAddressPrefix=$(BECH32_PREFIX)


BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


###########
# Install #
###########


all: install

.PHONY: install
install: build
	@echo "--> installing rollapp-evm"
	mv build/rollapp-evm $(GOPATH)/bin/rollapp-evm


.PHONY: build
build: go.sum ## Compiles the rollapd binary
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@echo "--> building rollapp-evm"
	@go build  -o build/rollapp-evm $(BUILD_FLAGS) ./cmd/rollappd


.PHONY: clean
clean: ## Clean temporary files
	go clean

clean-tmp: ## Clean temporary files
	rm -rf \
	tmp-swagger-gen/

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.7
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protogen.sh; fi
	@go mod tidy

#? proto-swagger-gen: Generate Protobuf Swagger
proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-clean:
	@echo "Cleaning proto generating docker container"
	@docker rm $(containerProtoGen) || true

.PHONY: proto-gen proto-swagger-gen proto-clean