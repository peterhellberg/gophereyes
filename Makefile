.PHONY: build

build:
	go build -ldflags="-s -w" -o build/gophereyes

install:
	go install -a -ldflags="-s -w"

run:
	go run *.go

clean:
	rm -rf build

app:
	macpack build
	mkdir -p build
	rm -rf build/GopherEyes.app
	mv GopherEyes.app build/

upx: build
	upx --brute build/gophereyes

upx-app: app
	upx --brute build/GopherEyes.app/Contents/MacOS/GopherEyes
