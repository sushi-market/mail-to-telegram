
[![Release](https://img.shields.io/github/v/release/petuhovskiy/mail-to-telegram)](https://github.com/petuhovskiy/mail-to-telegram/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/petuhovskiy/mail-to-telegram/master)](https://golang.org/doc/devel/release.html)
[![Docker](https://img.shields.io/docker/pulls/arthurwow/mailbot)](https://hub.docker.com/r/arthurwow/mailbot)

## About

`mail-to-telegram` listens to your mail (imap) and sends the message on telegram. The email updates are fetched with polling every 15 seconds.

## Install

### Docker

```bash
docker pull arthurwow/mailbot
docker run --rm arthurwow/mailbot --help
```

### docker-compose

docker-compose.yml
```yml
version: "3.7"
services:
  mailbot:
    image: arthurwow/mailbot
    env_file:
      - .env
    restart: always
```

.env
```.env
EMAIL_SERVER=imap.yandex.ru:993
EMAIL_LOGIN=your_email@yandex.ru
EMAIL_PASSWORD=your_password
TELEGRAM_TOKEN=123456789:AABB-telegram-bot-token-from-botfather
TELEGRAM_USER_ID=1234567894
```

To start:
```bash
docker-compose pull
docker-compose up -d
```