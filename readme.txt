> cd xendit-1
> touch Dockerfile
> touch docker-compose.yml
then, open project folder xendit-1
then, edit Dockerfile
-----------------------------------------------
FROM golang:1.19.0

WORKDIR /Users/macbook/Desktop/docker/xendit-1
-----------------------------------------------
then, edit docker-compose.yml
-----------------------------------------------
version: '3.8'

services:
  web:
    build: .
    ports:
      - 3000:3000
    volumes:
      - .:/Users/macbook/Desktop/docker/xendit-1
-----------------------------------------------
> docker compose up
> docker compose run --service-ports web bash
then, test the container:
# root@590dc64ac934:/Users/macbook/Desktop/docker/xendit-1# go version
# go version go1.19 linux/amd64
# root@5f74f7f4e6e3:/Users/macbook/Desktop/docker/xendit-1# go mod init github.com/jonipwi/xendit-1
# root@5f74f7f4e6e3:/Users/macbook/Desktop/docker/xendit-1# go get github.com/gofiber/fiber/v2
# root@5f74f7f4e6e3:/Users/macbook/Desktop/docker/xendit-1# mkdir cmd
# root@5f74f7f4e6e3:/Users/macbook/Desktop/docker/xendit-1# touch cmd/main.go
# root@5f74f7f4e6e3:/Users/macbook/Desktop/docker/xendit-1# go run cmd/main.go -b 0.0.0.0
then, edit Dockerfile:
-----------------------------------------------
FROM golang:1.19.0

WORKDIR /Users/macbook/Desktop/docker/xendit-1

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
-----------------------------------------------
then, edit docker-compose.yml:
-----------------------------------------------
version: '3.8'

services:
  web:
    build: .
    env_file:
      - .env
    ports:
      - 3000:3000
    volumes:
      - .:/Users/macbook/Desktop/docker/xendit-1
    command: air cmd/main.go -b 0.0.0.0
  db:
    image: postgres:alphine
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - 5432:5432
    volumes:
      - postgres-db:/Users/macbook/Desktop/docker/xendit-1/lib/postgresql/data

volumes:
  postgres-db:
-----------------------------------------------
> touch .air.toml
then, copy https://github.com/cosmtrek/air/blob/master/air_example.toml
# change line 10 cmd = "go build -o ./tmp/main ./cmd"
> touch .env
then, edit .env file:
-----------------------------------------------
DB_USER=root
DB_PASSWORD=root
DB_NAME=test
-----------------------------------------------
> docker compose build
> docker compose up
> docker compose run --service-ports web bash
# root@78c7882cb4e5:/Users/macbook/Desktop/docker/xendit-1# go get gorm.io/gorm
# root@78c7882cb4e5:/Users/macbook/Desktop/docker/xendit-1# go get gorm.io/driver/postgres
> docker pull postgres
> docker run -itd -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -v /Users/macbook/Desktop/docker/xendit-1/lib/postgresql/data --name postgresql postgres
5aeda2b20a708296d22db4451d0ca57e8d23acbfe337be0dc9b526a33b302cf5
> docker run -it --rm --network some-network postgres psql -h some-postgres -U postgres

------------------------ ACCESS TOKEN ------------
docker login -u jonipwi
password: dckr_pat_FJlbSp1Mp9-7HDMi5AAgv_F_WFc
docker build -t golang:v1 .
docker run -it --gpus all --privileged golang:v1
docker image push jonipwi/xendit-1-web:latest
