# scripts

MySQL

```bash
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=nuxbt -e TZ=Asia/Shanghai -p 5432:3306 -d mysql:8.4.1
```

Redis

```bash
docker run --name some-redis -d --health-cmd "redis-cli ping" --health-interval 5s --health-timeout 3s --health-retries 10 -p 6379:6379 redis
```

Minio

```bash
docker run --name some-minio -d -e MINIO_ACCESS_KEY=ChYm7ufIwNAOzq6PQPCA -e MINIO_SECRET_KEY=udicP52IwRbmo2hf6lFvjUS7NP5BhlAdsGNIuDE5 -e MINIO_DEFAULT_BUCKETS=nuxbt:public -p 9000:9000 bitnami/minio:latest
```
