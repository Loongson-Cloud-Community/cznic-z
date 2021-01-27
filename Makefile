# Copyright 2021 The Zlib-Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all bench clean cover cpu editor internalError later mem nuke todo edit devbench

grep=--include=*.go
ngrep='TODOOK\|internalError\|testdata\|scanner.go\|parser.go\|ast.go\|assets.go\|TODO-'


all:
	@LC_ALL=C date
	@go version 2>&1 | tee log
	@gofmt -l -s -w *.go
	@go install -v ./...
	@go test -i
	@go test 2>&1 -timeout 1h | tee -a log
	@go vet 2>&1 | grep -v $(ngrep) || true
	@golint 2>&1 | grep -v $(ngrep) || true
	@make todo
	@misspell *.go
	@staticcheck | grep -v 'scanner\.go' || true
	@maligned || true
	@grep -n --color=always 'FAIL\|PASS' log 
	LC_ALL=C date 2>&1 | tee -a log


darwin_amd64:
	TARGET_GOOS=darwin TARGET_GOARCH=amd64 go generate 2>&1 | tee /tmp/log-generate-zlib-darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -v
	gofmt -l -s -w *.go

linux_amd64:
	TARGET_GOOS=linux TARGET_GOARCH=amd64 go generate 2>&1 | tee /tmp/log-generate-zlib-linux-amd64
	GOOS=linux GOARCH=amd64 go build -v
	gofmt -l -s -w *.go

linux_386:
	CCGO_CPP=i686-linux-gnu-cpp GO_GENERATE_CC=i686-linux-gnu-gcc TARGET_GOOS=linux TARGET_GOARCH=386 go generate 2>&1 | tee /tmp/log-generate-zlib-linux-386
	GOOS=linux GOARCH=386 go build -v
	gofmt -l -s -w *.go

linux_arm:
	QEMU_LD_PREFIX=/usr/arm-linux-gnueabi CCGO_CPP=arm-linux-gnueabi-cpp-8 GO_GENERATE_CC=arm-linux-gnueabi-gcc-8 TARGET_GOOS=linux TARGET_GOARCH=arm go generate 2>&1 | tee /tmp/log-generate-zlib-linux-arm
	GOOS=linux GOARCH=arm go build -v
	gofmt -l -s -w *.go

linux_arm64:
	QEMU_LD_PREFIX=/usr/aarch64-linux-gnu CCGO_CPP=aarch64-linux-gnu-cpp-8 GO_GENERATE_CC=aarch64-linux-gnu-gcc-8 TARGET_GOOS=linux TARGET_GOARCH=arm64 go generate 2>&1 | tee /tmp/log-generate-zlib-linux-arm64
	GOOS=linux GOARCH=arm64 go build -v
	gofmt -l -s -w *.go

windows_amd64:
	#TODO CCGO_CPP=x86_64-w64-mingw32-cpp TARGET_GOOS=windows TARGET_GOARCH=amd64 go generate 2>&1 | tee /tmp/log-generate-zlib-windows-amd64
	#TODO GOOS=windows GOARCH=amd64 go build -v
	#TODO gofmt -l -s -w *.go

windows_386:
	#TODO CCGO_CPP=i686-w64-mingw32-cpp TARGET_GOOS=windows TARGET_GOARCH=386 go generate 2>&1 | tee /tmp/log-generate-zlib-windows-386
	#TODO GOOS=windows GOARCH=386 go build -v
	#TODO gofmt -l -s -w *.go

all_targets: linux_amd64 linux_386 linux_arm linux_arm64 windows_amd64 windows_386
	gofmt -l -s -w *.go
	echo done

devbench:
	date 2>&1 | tee log-devbench
	go test -timeout 24h -dev -run @ -bench . 2>&1 | tee -a log-devbench
	grep -n 'FAIL\|SKIP' log-devbench || true

bench:
	date 2>&1 | tee log-bench
	go test -timeout 24h -v -run '^[^E]' -bench . 2>&1 | tee -a log-bench
	grep -n 'FAIL\|SKIP' log-bench || true

clean:
	go clean
	rm -f *~ *.test *.out

cover:
	t=$(shell mktemp) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu: clean
	go test -run @ -bench . -cpuprofile cpu.out
	go tool pprof -lines *.test cpu.out

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	gofmt -l -s -w *.go
	nilness .
	GO111MODULE=off go install -v 2>&1 | tee log-install

later:
	@grep -n $(grep) LATER * || true
	@grep -n $(grep) MAYBE * || true

mem: clean
	go test -run @ -dev -bench . -memprofile mem.out -timeout 24h
	go tool pprof -lines -web -alloc_space *.test mem.out

nuke: clean
	go clean -i

todo:
	@grep -nr $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * | grep -v $(ngrep) || true
	@grep -nrw $(grep) 'TODO\|panic' * | grep -v $(ngrep) || true
	@grep -nr $(grep) BUG * | grep -v $(ngrep) || true
	@grep -nr $(grep) [^[:alpha:]]println * | grep -v $(ngrep) || true
	@grep -nir $(grep) 'work.*progress' || true
