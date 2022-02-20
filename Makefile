SHELL = /bin/bash

SERVICES = \
	backend/idempotency \
	backend/time/gateway \
	backend/time/service \

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
	@printf "\033[96m# 1st/2nd requests\033[0m\n"
	curl -i -XPOST -H "Idempotency-Key: 8e03978e-40d5-43e8-bc93-6894a57f9324" localhost:58081/v1/now&
	curl -i -XPOST -H "Idempotency-Key: 8e03978e-40d5-43e8-bc93-6894a57f9324" localhost:58081/v1/now

	@printf "\033[93m \n\n#   delay 3 sec \033[0m\n"
	sleep 3

	@printf "\033[96m \n\n# 3rd request\033[0m\n"
	curl -s -XPOST -H "Idempotency-Key: 8e03978e-40d5-43e8-bc93-6894a57f9324" localhost:58081/v1/now

	@printf "\033[96m \n\n# 4th request\033[0m\n"
	curl -s -XPOST -H "Idempotency-Key: 12345678-1234-4321-1234-123456789abc" localhost:58081/v1/now

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext localhost:50051 kzmake.time.v1.Time/Now
