# github.com/skingford/colly-spider

this is a colly spider demo

## docker

1. 交叉编译（mac）

```bash
sudo env GOOS=linux GOARCH=386 go build main.go
```

2. 编译镜像

```bash
# 版本号
docker build -t colly-spider:0.1 .

# 默认版本号
docker build -t colly-spider .
```

3. 启动 compose

```bash
docker-compose -f docker-compose.yml up -d colly
```
