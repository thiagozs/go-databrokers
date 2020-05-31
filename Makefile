GOCMD=go
GOBUILD=$(GOCMD) build
GOENV=$(GOCMD) env
OUTDIR=out
VERSION=1.0.0
LDFLAGS=-ldflags "-X main.version=${VERSION}"
NAME=server
