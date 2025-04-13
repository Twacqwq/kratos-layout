English | [简体中文](README_CN.md)

# Kratos Project Template
For personal use kratos-layout.

## Install Kratos
```bash
$ go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Using
```bash
$ kratos new myproject -r https://github.com/Twacqwq/kratos-layout
```

## Features
- [x] Rename `api` -> `proto`, Avoid multiple proto definitions resulting in multiple duplicate v1 packages
- [x] Update dependencies and use go 1.24
- [ ] ORM
- [ ] OTEL
- [ ] ... (welcome issue&pr)