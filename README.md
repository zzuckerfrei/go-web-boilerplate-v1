# go-web-poilerplate-v1

## Project Architecture
```
.
├── .github
│   └── workflows
│      └── local-ci.yaml
├── config
│   └── config.go
├── middleware
│   └── logger.go
├── Dockerfile 
├── go.mod
├── main.go
├── config-aks.yaml
├── config-eks.yaml
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

## Slack Notification
notification output sample
```
{
    "Repository": "Wondermove-Inc/w-lab-onprem-ci-test",
    "Ref": "refs/heads/main",
    "Tag Version": "v1.2.17",
    "Author": "flash-wondermove",
    "Status": "success",
    "Commit Message": "test commit message #onprem"
}
```