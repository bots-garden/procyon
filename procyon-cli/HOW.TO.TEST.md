

## Start Procyon Registry

```bash
REGISTRY_CRT="certs/procyon-registry.local.crt" \
REGISTRY_KEY="certs/procyon-registry.local.key" \
REGISTRY_HTTP=7070 \
REGISTRY_HTTPS=7070 \
REGISTRY_FUNCTIONS_PATH="./functions" \
node index.js 
```
> You need certificates

## Publish Wasm modules

```bash
go run main.go registry publish \
  --path ../samples/satellites/forty-two/forty-two.wasm \
  --function forty-two \
  --version 0.0.0

go run main.go registry publish \
  --path ../samples/satellites/hello-world-1.0.1/hello-world.wasm \
  --function hello-world \
  --version 1.0.1

go run main.go registry publish \
  --path ../samples/satellites/hello-world-1.0.2/hello-world.wasm \
  --function hello-world \
  --version 1.0.2
```

## Start Procyon Launcher

```bash
WASM_WORKER_PORT=9090 ./procyon-launcher
```

## Start Procyon Reverse

```bash
PROXY_HTTP=8080 ./procyon-reverse
```

## Deploy Wasm modules/functions

```bash
go run main.go functions deploy \
  --wasm hello-world.1.0.1.wasm \
  --function hello-world \
  --revision rev1

go run main.go functions deploy \
  --wasm hello-world.1.0.2.wasm \
  --function hello-world \
  --revision rev2

go run main.go functions deploy \
  --wasm forty-two.0.0.0.wasm \
  --function forty-two \
  --revision rev1
```

## Call functions

```bash
go run main.go functions call \
  --function hello-world \
  --revision rev1 \
  --method GET

go run main.go functions call \
  --function hello-world \
  --method GET

go run main.go functions revision \
	--function hello-world \
	--revision rev1 \
	--switch on

go run main.go functions revision \
	--function hello-world \
	--revision rev2 \
	--switch on

go run main.go functions call \
  --function hello-world \
  --revision rev2 \
  --method GET

go run main.go functions call \
  --function forty-two \
  --revision rev1 \
  --method GET
```