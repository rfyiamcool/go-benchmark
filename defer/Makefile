.DEFAULT_GOAL := all
.PHONY: boom

all: run

export GO15VENDOREXPERIMENT=1

run:
	go test -bench=. -run=none -benchmem .
