SHELL = /bin/bash

SERVICES = \
	backend/api/gateway \
	backend/svc/time \

.PHONY: all
all: pre proto fmt lint

.PHONY: install
install:
	make -C .devcontainer install


.PHONY: pre
pre:
	@for f in $(SERVICES); do make -C $$f pre; done

.PHONY: fmt
fmt:
	@for f in $(SERVICES); do make -C $$f fmt; done

.PHONY: lint
lint:
	@for f in $(SERVICES); do make -C $$f lint; done


.PHONY: proto
proto:
	cd api && buf mod update && cd -
	buf generate


.PHONY: build
build:
	skaffold build


.PHONY: kind
kind:
	kind get clusters -q | grep "_idempotencykey" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name _idempotencykey

.PHONY: dev
dev:
	skaffold dev


.PHONY: deploy-production
deploy-production:
	docker login ghcr.io
	skaffold run -p production

.PHONY: destroy-production
destroy-production:
	skaffold delete -p production

.PHONY: http
http:
	curl -i localhost:58080/v1/now

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext localhost:50001 kzmake.time.v1.Time/Now
