## Overview

This repository is a template for a Todo application in the Golang that can run on any platform as long as Docker is running.

## Endpoints

Method | Path                   | Description                                   |                                                                         
---    |------------------------|------------------------------------------------
GET    | `/health`              | Health page                                   |
GET    | `/v1/users/list`       | Получение пользователей системы по параметрам |
GET    | `/v1/users/get/{id}`   | Получение пользователя системы по id          |
POST   | `/v1/users/create`     | Создание нового пользователя                  |
PUT    | `/v1/users/update`     | Изменение данных пользователя                 |
DELETE | `/v1/users/delete/{id}`| Блокировка/удаление пользователя              |

## Usage

```shell
$ cp .env.example .env
```
