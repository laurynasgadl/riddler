GO     := go
GOFMT  := gofmt
GOPATH := $(shell $(GO) env GOPATH)

STATICCHECK := $(GOPATH)/bin/staticcheck
$(STATICCHECK):
	$(GO) install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: check
check: $(STATICCHECK)
	$(GOFMT) -s -w ./
	$(STATICCHECK) ./...
	$(GO) vet ./...
	$(GO) mod tidy

.PHONY: test
test:
	$(GO) test -v ./...
