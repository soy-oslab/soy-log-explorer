GOBIN=go
GORUN=$(GOBIN) run
GOBUILD=$(GOBIN) build

BINPATH=$(shell pwd)/bin

all: generator

clean:
	rm -rf $(BINPATH)/*

explorer:
	$(GORUN) ./cmd/explorer/main.go

explorer-build:
	mkdir -p $(BINPATH)/explorer
	$(GOBUILD) -o $(BINPATH)/explorer ./cmd/explorer/main.go

.PHONY: gen-src-archive
gen-src-archive:
	@echo "## explorer: make gen-src-archive"
	./scripts/explorer_gen_src_archive.sh
