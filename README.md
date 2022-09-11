# github.com/skingford/colly-spider

this is a colly spider demo

## docker

1. 编译镜像

```bash
# 版本号
docker build -t colly-spider:0.1 .

# 默认版本号
docker build -t colly-spider .

# 查看日志
docker logs -f -t --tail 10 docker_id
```

2. 启动 compose

```bash
docker-compose -f docker-compose.yml up -d server
```
