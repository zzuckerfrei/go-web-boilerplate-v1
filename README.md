# go-web-poilerplate-v1

## Project Architecture
```
.
├── .github
│   └── workflows
│      └── local-ci.yaml
├── config
│   ├── config-aks.yaml
│   ├── config-eks.yaml
│   └── config.go
├── middleware
│   └── logger.go
├── main.go
├── Dockerfile 
├── go.mod
└── test.log
```

## Version Management
* If commit message contains **"#onprem"** and **"#major"**, Image **major** version will be updated.
```
# example commit message
20240624 major update #onprem #major 

# Image version
v1.1.4 -> v2.0.0
```

* If commit message contains **"#onprem"** and **"#minor"**, Image **minor** version will be updated.
```
# example commit message
20240624 hot fix #onprem #minor

# Image version
v1.1.4 -> v1.2.0
```

* If commit message contains **"#onprem"**, Image **patch** version will be updated.
```
# example commit message
#onprem edit some-file.go

# Image version
v1.1.4 -> v1.1.5
```

## MatterMost Notification
notification output sample
```
{
    "Repository": "zzuckerfrei/go-web-boilerplate-v1",
    "Ref": "refs/heads/main",
    "Tag Version": "v0.1.0",
    "Author": "zzuckerfrei",
    "Status": "success",
    "Commit Message": "#onprem #minor test"
}
```