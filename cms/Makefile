export GO111MODULE := on
GO := go
GO_BUILD := $(GO) build -trimpath
GOBIN=$(shell go env GOPATH)/bin


gen-api: $(GOBIN)/goctl
	rm -rf internal/handler
	goctl api format -dir ./restapi
	goctl api go -api ./restapi/cms.api -dir . -style gozero -home ./.goctl
