## Overview

This repository is a template for a Todo application in the Golang that can run on any platform as long as Docker is running.

MySQL is used as the data store.

## Endpoints

Method | Path       | Description |                                                                         
---    |------------|--------------
GET    | `/health/` | Health page |

## Usage

```shell
$ cp .env.example .env

$ docker compose up -d --build
```

## Tips

### Connect db container

```shell
$ docker compose exec -it db /bin/bash -c "mysql -uroot -p<PASSWORD>"
```
