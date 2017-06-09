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
	rm -rf build/GopherEyes.app GopherEyes.app/Contents/Resources/css
	rm -f GopherEyes.app/Contents/Resources/.DS_Store
	rm -f GopherEyes.app/Contents/Resources/icon.png
	mv GopherEyes.app build/

upx: build
	upx --brute build/gophereyes

upx-app: app
	upx --brute build/GopherEyes.app/Contents/MacOS/GopherEyes
