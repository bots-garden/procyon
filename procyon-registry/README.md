# Procyon Registry

Small Wasm registry to provide Runnables (ansd other satellites) to Procyon launcher

## Launch Procyon Registry

> 🖐️ if you use **Sat** from **Suborbital**, you need to serve the wasm modules (Runnables for Sat) with https
```bash
REGISTRY_CRT="certs/procyon-registry.local.crt" \
REGISTRY_KEY="certs/procyon-registry.local.key" \
REGISTRY_HTTP=7070 \
REGISTRY_HTTPS=7070 \
REGISTRY_FUNCTIONS_PATH="./functions" \
node index.js 
```

If you need local certificates, follow the below procedure:

### Generate self-signed certificates

```bash
sudo apt-get update
sudo apt install libnss3-tools -y
brew install mkcert
```
> If you use Gitpod to open this project you don't need to install anything

> Generate the certificates for the Procyon Registry:
```bash
cd procyon-registry/certs
./generate.sh
```


## Upload wasm file

```bash
./procyon-registryctl publish ../samples/satellites/hello-world-1.0.1/hello-world.wasm hello-world 1.0.1

./procyon-registryctl publish ../samples/satellites/hello-world-1.0.2/hello-world.wasm hello-world 1.0.2

./procyon-registryctl publish ../samples/satellites/forty-two/forty-two.wasm forty-two 0.0.0
```

## Delete wasm file

```bash
🚧
```

## Download wasm file

```bash
./procyon-registryctl download hello-world.1.0.1.wasm

./procyon-registryctl download hello-world.1.0.2.wasm

./procyon-registryctl download forty-two.0.0.0.wasm
```


### Serve hello.wasm with Sat

```bash
SAT_HTTP_PORT=8080 ./executors/sat https://localhost:7070/get/forty-two.0.0.0.wasm

curl localhost:8080 -d 'Bob Morane'
```

