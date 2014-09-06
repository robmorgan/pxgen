build: vet test compile

test:
	go test
	
vet:
	go vet ./...

compile:
	@go get -ldflags ./...

size:
	du -hs $$(which pxgen)
