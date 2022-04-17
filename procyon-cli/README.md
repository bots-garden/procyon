# Procyon CLI

```bash
go mod init github.com/bots-garden/procyon/procyon-cli
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

cobra-cli init \
--author "Philippe Charrière @k33g_org" \
--license MIT \
--viper
```


> Add a command
```bash
cobra-cli add hello

```



## Commands

### Registry

> several keywords
```bash
cobra-cli add registry
cobra-cli add url -p 'registryCmd'
cobra-cli add publish -p 'registryCmd'
cobra-cli add download -p 'registryCmd'

# only for the procyon registry
procyon-cli registry url
procyon-cli registry publish <path_to_wasm_file> <function_name(*)> <version>
*: or service_name
procyon-cli registry download <path_to_wasm_file>
```


### Deployment (from a registry)


```bash
# old version
> WASM_REGISTRY_URL="https://localhost:7070/get"
procyonctl task deploy hello-world.1.0.1.wasm hello-world rev1
> WASM_REGISTRY_URL="https://registry-cdn.wapm.io/contents"
procyonctl task deploy k33g/hello-world/1.0.1/hello-world.wasm hello-world rev1

```