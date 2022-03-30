# Procyon

[![Open in GitPod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/bots-garden/procyon)


## 🚧 This is a work in progress


```bash
curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hello.wasm",
      "wasmFunctionHttpPort": 8082,
      "wasmRegistryUrl": "https://localhost:9999/hello/hello.wasm"
    }
  ' http://localhost:9090/tasks
```

```bash
curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hey.wasm",
      "wasmFunctionHttpPort": 8083,
      "wasmRegistryUrl": "https://localhost:9999/hey/hey.wasm"
    }
  ' http://localhost:9090/tasks
```

```bash
curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "",
      "wasmFunctionHttpPort": 8084,
      "wasmRegistryUrl": "https://localhost:9999/hi/hi.wasm"
    }
  ' http://localhost:9090/tasks
```

```bash
curl -v http://localhost:9090/tasks
```

```bash
curl -v --request DELETE \
http://localhost:9090/tasks/9babdc47-1df6-4b03-9133-b91e4380598c
```


## What send to the API to create a task?

- To start a task, the state of the task has to be "Scheduled" (`1`)
- To stop a task, the state of the task has to be "Completed" (`3`)

```json
{
  "executor": 1, // 1: galago 2:sat
  "wasmFileName": "hello.wasm",
  "wasmFunctionHttpPort": 8082,
  "wasmRegistryUrl": "https://localhost:9999/hello/hello.wasm"
}
```

**States**:
```
- Pending   0
- Scheduled 1
- Running   2
- Completed 3
- Failed    4
```