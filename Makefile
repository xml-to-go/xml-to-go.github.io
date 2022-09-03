build:
	npm i
	npm run build

webpack:
	npm run build
	# time --format="webpack took %E" npm run build

esbuild:
	npm run esbuild
	#time --format="esbuild took %E" npm run esbuild

browse:
	browse index.html

wasm:
	# go mod edit -go=1.19
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./src/wasm/
	GOOS=js GOARCH=wasm go build -o ./static/js/wasm/xml-to-go.wasm ./go/wasm/main.go
