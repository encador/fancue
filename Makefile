.PHONY: run dev clean build

clean:
	- rm -rf bin

build:
	go tool templ generate
	go build -o bin/fancue ./

run: build
	./bin/fancue

dev: build
	go tool templ generate --watch \
		--cmd="./bin/fancue" \
		--proxy="http://localhost:55000" \
		--proxybind="localhost" --proxyport="8080" \
		--open-browser=false
