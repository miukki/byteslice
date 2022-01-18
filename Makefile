#!/usr/bin/env bash

.DEFAULT_GOAL := help
.PHONY: test lint install-linters
DIR = $(shell pwd)

test:
	which gotestsum || ( \
		go get gotest.tools/gotestsum@v1.6.4 \
	)
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" gotestsum --format testname -- -timeout=5m --count=1 -test.short ./...


lint: install-linters ## Run linters.
	golangci-lint run --deadline=3m --config .golangci.yml ./...

install-linters: ## Install linters
	go get mvdan.cc/gofumpt@v0.1.1
	go get golang.org/x/tools/cmd/goimports@v0.1.1
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
	go get gotest.tools/gotestsum@v1.6.4

install_cmd:
	GOBIN=${DIR}/bin go install ./cmd/...

docker_run:
	docker run --env-file ./scripts/env-docker.sh -it -p 3003:3003 testzkp

docker_build:
	docker build -t testzkp --build-arg "GIT_COMMIT=${GIT_COMMIT}" -f Dockerfile .

vendor: install-linters
	go get github.com/nomad-software/vend
	vend

generate:
	go generate cmd/poc/byteslice/circuit/circuit.go

circuit:
	go run cmd/poc/main.go circuit

sha3:
	go run cmd/poc/main.go sha3

sha3_apikey:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --apikey

sha3_deploy:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --deploy --zkpkey=${ZKP_KEY}

sha3_loghash:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --loghash --address=${ADDR} --txn=${TXN} --type=${TTYPE} --state=${STATE} --zkpkey=${ZKP_KEY}

sha3_getlog:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --getlog --address=${ADDR} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_gethash:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --gethash --address=${ADDR} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_gethistory:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --gethistory --address=${ADDR} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_verifylatest:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --verifylatest --address=${ADDR} --txn=${TXN} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_deploy_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --deploy --smilo --zkpkey=${ZKP_KEY}

sha3_loghash_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --loghash --smilo --address=${ADDR} --txn=${TXN} --type=${TTYPE} --state=${STATE} --zkpkey=${ZKP_KEY}

sha3_getlog_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --getlog --smilo --address=${ADDR} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_gethash_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --gethash --smilo --address=${ADDR} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_gethistory_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --gethistory --smilo --address=${ADDR} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

sha3_verifylatest_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go sha3 --verifylatest --smilo --address=${ADDR} --txn=${TXN} --id=${ID} --type=${TTYPE} --zkpkey=${ZKP_KEY}

byteslice_init:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go byteslice --init

byteslice_chain:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go byteslice --chain

byteslice_debug:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go byteslice --debug

byteslice:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go byteslice

server:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go server

signing:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go signing  --msg=${MSG} --zkpkey=${ZKP_KEY}

signing_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go signing --smilo  --msg=${MSG} --zkpkey=${ZKP_KEY}

verifysignature:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go signing --verifysig --msg=${MSG} --zkpkey=${ZKP_KEY} --pubkey=${PUBKEY} --signature=${SIGNATURE}

verifyfilesignature:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go signing --verifyfilesig --file=${FILE} --zkpkey=${ZKP_KEY} --pubkey=${PUBKEY} --signature=${SIGNATURE}

signing_file:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go signing  --filesign --file=${FILE} --zkpkey=${ZKP_KEY}

signing_file_smilo:
	GOBIN=$(PWD)/bin PATH="${HOME}/.local/bin:${PATH}" go run cmd/poc/main.go signing --smilo  --filesign --file=${FILE} --zkpkey=${ZKP_KEY}


