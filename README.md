### Телеграмм бот для получения текущей погоды. Работает на основе данных OpenWeatherMap

## Deployment

### docker compose

**Build** application

```shell
docker compose build
```

**Run** application

```shell
docker compose up -d
```

### All options are loaded from **[.env](.env)**

```dotenv
BOT_DEBUG=false
BOT_TELEGRAM_TOKEN=TOKEN

OPENWEATHERMAP_API_KEY=API_KEY
```