## Overview

This repository is a main service of loyalty system appliaction.

## Endpoints

Method | Path                   | Description                                   |                                                                         
---    |------------------------|------------------------------------------------
GET    | `/health`              | Health page                                   |
GET    | `/metrics`             | Страница с метриками                          |
---    |------------------------|-----------------------------------------------|
GET    | `/v1/users/list`       | Получение пользователей системы по параметрам |
GET    | `/v1/users/get/{id}`   | Получение пользователя системы по id          |
POST   | `/v1/users/create`     | Создание нового пользователя                  |
PUT    | `/v1/users/update`     | Изменение данных пользователя                 |
DELETE | `/v1/users/delete/{id}`| Блокировка/удаление пользователя              |
GET    | `/v1/users/me`         | Получение информации по своему пользователю   |
PUT    | `/v1/users/me`         | Изменение информации своего пользователю      |
PUT    | `/v1/users/me/deposit` | Внесение депозита своему пользователю         |
---    |------------------------|-----------------------------------------------|
GET    | `/v1/items/list`       | Получение списка товаров                      |
POST   | `/v1/items/buy/{id}`   | Покупка данного товара текущим пользователем  |

## Usage

```shell
$ cp .env.example .env
cd src
go run
```
