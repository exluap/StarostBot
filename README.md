# StarostBot [![Build Status](http://drone.exluap.xyz/api/badges/exluap/StarostBot/status.svg)(http://drone.exluap.xyz/exluap/StarostBot)
Бот для уведомлений определенных сообщений ВК


## Фукнционал

* Фильтрация сообщений по команде /all
* Оповещение всех пользователей подписанных на оповещения

## Настройка 

* Скачайте исходники бота с помощью `git clone https://github.com/exluap/starostbot`
* В файле `main.go` замените `DIALOG_ID` (46 строка) на ID беседы в ВК следующего формата `2000000000+id` чтобы было, например `2000000009` где 9 - id диалога 
* Получите VK Access Token по ссылке `https://oauth.vk.com/authorize?client_id=6058624&scope=messages,offline&redirect_uri=http://api.vk.com/blank.html&display=page&response_type=token`
* Скопируйте Access Token полученный в ссылке
* Создайте файл config.json следующего содержания: 
```
{
  "VK": {
    "key": "ACCESS_TOKE"
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
