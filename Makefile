# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

GO_VERSION=1.14.3
DEV_DIR=../

test: fmt
	@go test -v

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@cd $(DEV_DIR) ; wget https://dl.google.com/go/go$(GO_VERSION).linux-amd64.tar.gz
	@cd $(DEV_DIR) ; tar xvzf go$(GO_VERSION).linux-amd64.tar.gz

install:
	go get -u github.com/chai2010/gettext-go
