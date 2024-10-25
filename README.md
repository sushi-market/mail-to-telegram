
## About

`mail-to-telegram` listens to your mail (imap) and sends the message on telegram. The email updates are fetched with polling every 15 seconds.

## Install

### Docker

```bash
docker build -t df/mailbot:v1 .
```

.env
```.env
EMAIL_SERVER=imap.yandex.ru:993
EMAIL_LOGIN=your_email@yandex.ru
EMAIL_PASSWORD=your_password
TELEGRAM_TOKEN=123456789:AABB-telegram-bot-token-from-botfather
TELEGRAM_USER_ID=1234567894
READ_TIMEOUT=15
```

To start:
```bash
docker compose up -d
```
