all:
	docker build --tag=market .
	docker run --interactive market

.PHONY: all