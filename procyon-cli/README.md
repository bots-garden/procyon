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
```


### Deployment (from a registry)

```bash
cobra-cli add services
cobra-cli add deploy -p 'servicesCmd'
cobra-cli add revision -p 'servicesCmd' (+switch)
cobra-cli add call -p 'servicesCmd'
  === flag ? ===
  cobra-cli add post -p 'servicesCmd'
  cobra-cli add get -p 'servicesCmd'
cobra-cli add list -p 'servicesCmd'
```

```bash
- procyonctl task deploy k33g/forty-two/1.0.0/forty-two.wasm forty-two rev1
  procyonctl task deploy k33g/hello-world/1.0.1/hello-world.wasm hello-world rev1

procyonctl task set-default-revision forty-two rev1 on

#procyonctl task set-default-revision hello-world rev2 on

procyonctl task set-default-revision hello-world rev1 on
procyonctl task set-default-revision hello-world rev2 on


procyonctl task set-default-revision hello-world rev2 off

========
🤔
- procyonctl task switch-func-revision forty-two rev1
- procyonctl task switch-func-revision hello-world rev1
========

- procyonctl func post hello 'Jane Doe'
- procyonctl func post-revision hello rev1 'John Doe'

- procyonctl func get forty-two
- procyonctl func get-revision hello-world rev1

- procyonctl func list
- procyonctl func help
```



```bash
# old version
> WASM_REGISTRY_URL="https://localhost:7070/get"
procyonctl task deploy hello-world.1.0.1.wasm hello-world rev1
> WASM_REGISTRY_URL="https://registry-cdn.wapm.io/contents"
procyonctl task deploy k33g/hello-world/1.0.1/hello-world.wasm hello-world rev1
```

### Task

cobra-cli add tasks
cobra-cli add all -p 'tasksCmd'
cobra-cli add info -p 'tasksCmd'
cobra-cli add kill -p 'tasksCmd'


```bash
- procyonctl task list
- procyonctl task info a6419ab3-fbe4-4f5e-92ea-c832d8869090
- procyonctl task switch-revision a6419ab3-fbe4-4f5e-92ea-c832d8869090
- procyonctl task kill a6419ab3-fbe4-4f5e-92ea-c832d8869090
  TODO: update task list when remove
- procyonctl task help
```