GO_VER ?= go
SRCDIR ?= src
GOOS ?= linux 
GOARCH ?= amd64 
TYPE ?= module

default: build

build: 
	for type in $(TYPE) ; do \
		for action in $(SRCDIR)/$$type/* ; do \
			GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_VER) build -o build/$$action/main $$action/main.go ;\
        done; \
	done

zip:
	startingdirectory=`pwd` ; \
 	for type in $(TYPE) ; do \
		for action in build/src/$$type/* ; do \
			mkdir -p ./iac/$$action ; \
			cd $$action ; zip ../../../../iac/$$action/main.zip main; cd $$startingdirectory ; \
        done; \
	done

tfyolo: 
	cd iac; terraform apply --auto-approve; cd ../

deploy: build set-file-time zip tfyolo

set-file-time: 
	find . -exec touch -t `git ls-files -z . | \
	xargs -0 -n1 -I{} -- git log -1 --date=format:"%Y%m%d%H%M" --format="%ad" {} | \
	sort -r | head -n 1` {} +

.PHONY: \
	build \
	zip \
	tfyolo \
	deploy \
	set-file-time \