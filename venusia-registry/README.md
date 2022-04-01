# Venusia

Small Wasm registry

## Upload wasm file

```bash
curl -v \
  -F "file=@oh.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload
```

## Delete wasm file

```bash
curl -v \
	-X DELETE https://localhost:9999/wasm/delete/oh.wasm
```

## Download wasm file

```bash
curl https://localhost:9999/wasm/download/oh.wasm --output oh.wasm
```

## Build Venusia


```bash
go build
```

## Create a Runnable

```bash
cd samples
subo create runnable hello --lang tinygo
subo build hello/
# it will generate a hello.wasm file
```

## Launch Venusia

```bash
REGISTRY_CRT="certs/venusia.local.crt" \
REGISTRY_KEY="certs/venusia.local.key" \
REGISTRY_HTTP=7070 \
./venusia
```

### Publish a wasm file to Venusia

```bash
curl -v \
  -F "file=@samples/hello/hello.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:7070/wasm/upload
```
> it will upload the file to the `functions` directory

### Serve hello.wasm with Sat

```bash
SAT_HTTP_PORT=8080 ./sat/.bin/sat https://localhost:7070/wasm/download/hello.wasm

curl localhost:8080 -d 'Bob Morane'
```

## Generate self-signed certificates

```bash
sudo apt-get update
sudo apt install libnss3-tools -y
brew install mkcert
```
> If you use Gitpod to open this project you don't need to install anything
