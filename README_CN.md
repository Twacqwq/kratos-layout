# Kratos Project Template
自用 kratos-layout 模板

## 安装
```bash
$ go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## 使用模板
```bash
$ kratos new myproject -r https://github.com/Twacqwq/kratos-layout
$ cd myproject
$ make all # 必须执行一次
```

## 特性
- [x] 重命名 `api` -> `proto`, 解决当有多个服务定义时产生的多个 v1 包命名, 降低维护的难度
- [x] 升级 go 版本为 1.24 并更新相关依赖
- [ ] ORM
- [ ] OTEL
- [ ] ...(welcome issue&pr)
