# StarostBot
Бот для старост МИРЭА


## Фукнционал

* Фильтрация сообщений по команде /all
* Оповещение всех пользователей подписанных на оповещения

## Настройка 

* Скачайте исходники бота с помощью `git clone https://github.com/exluap/starostbot`
* В файле `main.go` замените `DIALOG_ID` (46 строка) на ID беседы в ВК следующего формата `2000000000+id` чтобы было, например `2000000009` где 9 - id диалога
* Создайте файл config.json следующего содержания: 
```
{
  "VK": {
    "key": "ACCESS_TOKE",
    "dialog": DIALOG_ID
  },
  "OneSignal": {
    "APP_KEY": "ONESIGNAL_APP_ID",
    "USER_KEY": "ONESIGNAL_USER_KEY",
    "REST_API_KEY": "ONESIGNAL_APP_REST_KEY"
  }
}
```

* Скомпилируйте код с помощью GO
* Пользуйтесь
