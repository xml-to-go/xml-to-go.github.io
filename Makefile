webpack:
	npm run build

esbuild:
	npm run esbuild

# go mod edit -go=1.20
wasm:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./src/wasm/
	GOOS=js GOARCH=wasm go build -o ./static/js/wasm/xml-to-go.wasm ./go/wasm/*.go

server:
	go run ./go/server

browse:
	browse index.html

firefox-browse:
	firefox index.html
