# StarostBot [![Build Status](http://drone.exluap.xyz/api/badges/exluap/StarostBot/status.svg)](http://drone.exluap.xyz/exluap/StarostBot)
Бот для уведомлений определенных сообщений ВК


## Фукнционал

* Фильтрация сообщений по определенному хэштегу или слову
* Оповещение всех пользователей подписанных на оповещения [OneSignal](https://onesignal.com)

## Настройка и установка

* Скачайте нужную вам сборку под ваше устройстройство, под которым вы будете работать [отсюда](https://github.com/exluap/StarostBot/releases)
* Запустите бота
* Дождитесь скачивания ботом файла `config.example.json`
* Измените данные в файле `config.example.json`
* Измените `Key` в блоке `VK` на свой access token полученный по [ссылке](https://oauth.vk.com/authorize?client_id=6058624&scope=messages,offline&redirect_uri=http://api.vk.com/blank.html&display=page&response_type=token)
* Пользуйтесь

## ToDo

- [ ] Автопоиск Access Token по ссылке от пользователя
- [ ] Автонастройка с помощью мастера настройки
