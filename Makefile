
export PACKAGE := $(abspath .)
export OUTDIR := $(abspath ./bin)

export GOBIN := ${OUTDIR}:${GOPATH}/bin:${GOBIN}
export PATH := ${PATH}:${GOBIN}
export TMPDIR := /tmp

export GIT_SSH_COMMAND='ssh -o ControlMaster=no'

# Use GOPATH if provided OR assume default GOPATH - WORKS ON MAC & LINUX
ifndef GOPATH
	override GOPATH = '${HOME}/go'
endif

all: build run

sol:
	go get -u -v github.com/ethereum/go-ethereum
	make -C ${GOPATH}/src/github.com/ethereum/go-ethereum/
	make -C ${GOPATH}/src/github.com/ethereum/go-ethereum/ devtools

	solc --abi --bin --overwrite SmartTgStats.sol -o ${TMPDIR}

	abigen --bin ${TMPDIR}/SmartTgStats.bin --abi ${TMPDIR}/SmartTgStats.abi --pkg main --type SmartTgStats --out SmartTgStats.go

build:
	go get -d .
	go build -o ${OUTDIR}/run_backend

run:
	${OUTDIR}/run_backend
