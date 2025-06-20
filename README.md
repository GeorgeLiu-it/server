# GO application deployment using docker

## System setup

Refert to README_env.md

## Docker images

```bash
# Export
docker save -o [image.tar] [local-image:version]

# Import
docker load < [image.tar]
```

## Application deployment 

### Elasticsearch

**Docker command**

```bash
docker run -d --name blog_elasticsearch \
  -p 9200:9200 \
  -p 9300:9300 \
  -v $HOME/elasticsearch/data:/usr/share/elasticsearch/data \
  -v $HOME/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml:ro \
  -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
  -e "ELASTIC_PASSWORD=123456" \
  docker.elastic.co/elasticsearch/elasticsearch:9.0.0

# Elastic yaml config
# cat elasticsearch.yml
cluster.name: blog-cluster
node.name: single-node
network.host: 0.0.0.0
http.port: 9200
discovery.type: single-node
xpack.security.enabled: false

# Test es connection
curl -X GET http://192.168.68.231:9200
curl -u elastc:123456 http://192.168.68.231:9200
```

### Redis

**Docker command**

```bash
# Redis with local persistence
docker run -d --name blog_redis\
  -p 6379:6379 \
  -v $HOME/redis/data:/data \
  redis:latest \
  redis-server --appendonly yes --appendfsync everysec
```

### Mysql

**Docker command**

```bash
docker run -d \
  --name blog_mysql \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -e MYSQL_DATABASE=blog_db \
  -e MYSQL_USER=george \
  -e MYSQL_PASSWORD=123456 \
  -p 3306:3306 \
  -v $HOME/mysql/data:/var/lib/mysql \
  --restart unless-stopped \
  mysql:8
```

### Go application

**Code level initialization**

```bash
# cd to server folder
go mod tidy

# initiate db, es, admin
go run main.go -sql
go run main.go -avaya
go run main.go -es
go run main.go -admin
```

**Build application as binary**

- Window environment

  ```cmd
  # Open cmd, and cd to root folder
  set GOOS=linux
  set GOARCH=amd64
  go mod tidy
  go build main.go
  ```

- Linux

  ```bash
  go mod tidy
  go build main.go
  ```

**Build docker image**

* Dockerfile

  ```dockerfile
  # ---- Build stage ----
  FROM golang:1.23.2-alpine AS builder
  
  # Add your private CA
  COPY Zscaler_Root_CA.crt /usr/local/share/ca-certificates/ca.crt
  RUN update-ca-certificates
  
  # Add git + ca-certificates (required for HTTPS + proxy trust)
  RUN apk add --no-cache git ca-certificates
  
  # Set working dir
  WORKDIR /app
  
  COPY go.mod go.sum ./
  RUN go mod download
  
  # Copy source code
  COPY . .
  
  # Build binary
  RUN CGO_ENABLED=0 GOOS=linux go build -o app
  
  # ---- Final stage ----
  FROM scratch
  
  # Copy binary
  COPY --from=builder /app/app /app
  
  # Copy config
  COPY config.yaml /config.yaml
  
  # Copy CA certs (including your private CA)
  COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
  
  ENTRYPOINT ["/app"]
  ```

* Run 

  ```bash
  # Generate docker image
  docker build -t blog_server .
  
  # Run docker container
  docker run -d -p 8080:8080 -v $HOME/app/config.yaml:/config.yaml --name blog_server blog_server:latest
  ```

### Docker-compose to run all-in-0ne

**.env file**

```txt
# Paths
HOME_DIR=/home/george

# ES
ES_JAVA_OPTS=-Xms1g -Xmx1g
ELASTIC_PASSWORD=123456

# Mysql
MYSQL_ROOT_PASSWORD=123456
MYSQL_DATABASE=blog_db
MYSQL_USER=george
MYSQL_PASSWORD=123456
```

**docker-compose.yml**

```yml
version: '3.8'

services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:9.0.0
    container_name: blog_elasticsearch
    ports:
      - "9200:9200"
      - "9300:9300"
    environment:
      - ES_JAVA_OPTS=${ES_JAVA_OPTS}
      - ELASTIC_PASSWORD=${ELASTIC_PASSWORD}
    volumes:
      - ${HOME_DIR}/elasticsearch/data:/usr/share/elasticsearch/data
      - ${HOME_DIR}/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml:ro
    restart: unless-stopped
    networks:
      - blog_net

  redis:
    image: redis:latest
    container_name: blog_redis
    ports:
      - "6379:6379"
    volumes:
      - ${HOME_DIR}/redis/data:/data
    command: ["redis-server", "--appendonly", "yes", "--appendfsync", "everysec"]
    restart: unless-stopped
    networks:
      - blog_net

  mysql:
    image: mysql:8
    container_name: blog_mysql
    ports:
      - "3306:3306"
    environment:
      - MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
      - MYSQL_DATABASE=${MYSQL_DATABASE}
      - MYSQL_USER=${MYSQL_USER}
      - MYSQL_PASSWORD=${MYSQL_PASSWORD}
    volumes:
      - ${HOME_DIR}/mysql/data:/var/lib/mysql
    restart: unless-stopped
    networks:
      - blog_net

  blog_server:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: blog_server
    ports:
      - "8080:8080"
    volumes:
      - ${HOME_DIR}/app/config.yaml:/config.yaml:ro
    depends_on:
      - elasticsearch
      - redis
      - mysql
    networks:
      - blog_net
    healthcheck:
      test: ["CMD", "curl", "-f", "http://127.0.0.1:8080/api/base/health"]
      interval: 30s
      timeout: 10s
      retries: 3
    restart: unless-stopped

networks:
  blog_net:
    driver: bridge
```

**Docker-compose command**

```bash
docker compose up --build -d
```

## Host deployment

### Pull from github

```bash
# pull source code
git clone -b develop --single-branch https://github.com/GeorgeLiu-it/server.git
```

**Add .env**

```txt
# Paths
HOME_DIR=/home/george

# ES
ES_JAVA_OPTS=-Xms1g -Xmx1g
ELASTIC_PASSWORD=123456

# Mysql
MYSQL_ROOT_PASSWORD=123456
MYSQL_DATABASE=blog_db
MYSQL_USER=george
MYSQL_PASSWORD=123456
```

**Run docker compose**

```bash
# cd to server folder
docker compose build
docker compose up -d

# stop container
docker compose stop

# remove container
docker compose down
```
