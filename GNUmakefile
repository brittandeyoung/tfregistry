GO_VER ?= go
SRCDIR ?= src/api
GOOS ?= linux
GOARCH ?= arm64
TYPE ?= namespace

default: build

test:
	$(GO_VER) test ./...

build:
	for type in $(TYPE) ; do \
		for action in $(SRCDIR)/$$type/* ; do \
			GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_VER) build -o build/$$action/bootstrap $$action/main.go ;\
        done; \
	done

zip:
	startingdirectory=`pwd` ; \
 	for type in $(TYPE) ; do \
		for action in build/src/api/$$type/* ; do \
			mkdir -p ./iac/$$action ; \
			cd $$action ; zip ../../../../../iac/$$action/bootstrap.zip bootstrap; cd $$startingdirectory ; \
        done; \
	done

tfyolo:
	cd iac; terraform apply --auto-approve; cd ../

tfyodo:
	cd iac; terraform apply -destroy --auto-approve ; cd ../

deploy: build set-file-time zip tfyolo

set-file-time:
	find . -exec touch -t `git ls-files -z . | \
	xargs -0 -n1 -I{} -- git log -1 --date=format:"%Y%m%d%H%M" --format="%ad" {} | \
	sort -r | head -n 1` {} +

clean:
	rm -rf ./iac/build ./build

.PHONY: \
	build \
	zip \
	tfyolo \
	tfyodo \
	deploy \
	set-file-time \
	clean
