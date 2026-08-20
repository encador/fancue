.PHONY: run dev clean build

clean:
	- rm -rf bin

build:
	go tool templ generate
	go build -o bin/fancue ./

run: build
	./bin/fancue

dev:
	go tool templ generate --watch \
		--cmd="go run . -port=55002" \
		--proxy="http://localhost:55002" \
		--proxybind="localhost" --proxyport="8080" \
		--open-browser=false
