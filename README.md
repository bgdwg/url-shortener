# url-shortener

Минималистичный сервис, предоставляющий API по созданию сокращённых ссылок.

В качестве хранилища используется PostgreSQL.

# Как использовать

Чтобы собрать и запустить сервис, достаточно в директории с кодом запустить команду:

```sh
docker-compose up --build
```

Также при необходимости настроить переменные окружения в файле `docker-compose.yaml`.

Ссылка на Docker-образ: https://hub.docker.com/repository/docker/bgdwg/url-shortener-service
