

## Start Procyon Registry

```bash
cd procyon-registry
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
  --version 0.0.0 \
  --config .procyon-cli.yaml

go run main.go registry publish \
  --path ../samples/satellites/hello-world-1.0.1/hello-world.wasm \
  --function hello-world \
  --version 1.0.1 \
  --config .procyon-cli.yaml

go run main.go registry publish \
  --path ../samples/satellites/hello-world-1.0.2/hello-world.wasm \
  --function hello-world \
  --version 1.0.2 \
  --config .procyon-cli.yaml
```

## Start Procyon Launcher

```bash
cd procyon-launcher
PROCYON_WASM_WORKER_PORT=9090 PROCYON_ADMIN_TOKEN="ilovepandas" ./procyon-launcher
```

## Start Procyon Reverse

```bash
cd procyon-reverse-proxy
PROXY_HTTP=8080 ./procyon-reverse
```

## Deploy Wasm modules/functions

```bash
go run main.go functions deploy \
  --wasm https://localhost:7070/get/hello-world.1.0.1.wasm \
  --function hello-world \
  --revision rev1 \
  --config .procyon-cli.yaml

go run main.go functions deploy \
  --wasm https://localhost:7070/get/hello-world.1.0.2.wasm \
  --function hello-world \
  --revision rev2 \
  --config .procyon-cli.yaml

go run main.go functions deploy \
  --wasm https://localhost:7070/get/forty-two.0.0.0.wasm \
  --function forty-two \
  --revision rev1 \
  --config .procyon-cli.yaml
```

## Call functions

```bash
go run main.go functions call \
  --function hello-world \
  --revision rev1 \
  --method GET \
  --config .procyon-cli.yaml

go run main.go functions call \
  --function hello-world \
  --method GET \
  --config .procyon-cli.yaml

go run main.go functions revision \
	--function hello-world \
	--revision rev1 \
	--switch on \
  --config .procyon-cli.yaml

go run main.go functions revision \
	--function hello-world \
	--revision rev2 \
	--switch on \
  --config .procyon-cli.yaml

go run main.go functions call \
  --function hello-world \
  --revision rev2 \
  --method GET \
  --config .procyon-cli.yaml

go run main.go functions call \
  --function forty-two \
  --revision rev1 \
  --method GET \
  --config .procyon-cli.yaml
```

## Functions List

```bash
go run main.go functions list \
  --config .procyon-cli.yaml
```

## Tasks commands

```bash
go run main.go tasks list \
  --config .procyon-cli.yaml
```

```bash
go run main.go tasks info --task-id 1baa7939-a698-4112-b14d-fb9b35d0fac1 \
  --config .procyon-cli.yaml
```

```bash
go run main.go tasks kill --task-id 1baa7939-a698-4112-b14d-fb9b35d0fac1 \
  --config .procyon-cli.yaml
```

