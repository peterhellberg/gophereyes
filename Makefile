build:
	go build -ldflags="-s -w"

install:
	go install -a -ldflags="-s -w"

run:
	go run *.go

clean:
	rm -f gophereyes
