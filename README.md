# Тестовое задание

## Исходные данные:
### 1. Программа минимум
- 2 http роутера на разных портах
- На 1 порту метод set, на другом порту метод get.
  set - post запрос, в body запроса json модели UserGrade. Сохраняем структуру в storage
  get - get запрос, параметр в urlencoded ?user_id=, на выходе json модели UserGrade из storage
- Реализовать пакет storage. Задача пакета - хранить в рам структуру UserGrade по стринговому ключу. Ключ UserId
  Имеет публичные методы set, get
- Модель UserGrade
```go
  type UserGrade struct {
    UserId        string json:"user_id" validate:"required"
    PostpaidLimit int    json:"postpaid_limit"
    Spp           int    json:"spp"
    ShippingFee   int    json:"shipping_fee"
    ReturnFee     int    json:"return_fee"
  }
 ```
- Реализовать middleware с basic auth, закрыть им метод set
- В set могут присылать данные частями. Одним запросом Spp, след ShippingFee и т.д.

## 2. Программа максимум (hard skill)
- Реализовать репликацию в методе set. Для репликации используем брокер сообщений (nats streaming/kafka на ваш вкус - поднимайте локлько в докере)
  При получении данных в метод set, сервис публикует сообщение в канал. В горутине сервис подписывается на этот же канал.
  Отфильтровывает свои сообщения и обрабатывает сообщения других реплик.
- Реализовать метод /backup. Метод при запросе генерит дамп файл локальных данных в формате csv.gz, передает в response
  В бекап зашить время, когда он был сгенерирован.
- При старте приложения мы дергаем метод /backup реплики. Заполняем данными storage.
  Подписываемся к каналу с того времени, которое указано в бекапе.
- Подумайте над порядком запуска функций, бекап может быть большим и восстановление может занять время

## Описание решения
Структура проекта была сделана в соответствии со следующим [шаблоном](https://github.com/golang-standards/project-layout)

- Программа минмимум реализована полностью.
- В качестве брокера сообщений был выбран nats streaming.
- Реализована репликация в методе set. Сервис подписан на канал в горутине.
При получении собщений происходит вывод их в консоль.
- Также реализован метод /backup. Метод при запросе генерит дамп файл локальных данных в формате csv.gz,
передает в response. В бекап зашито время, когда он был сгенерирован (в имя файла).
- При старте приложения мы дергаем метод /backup реплики. Заполняем данными storage.
  Подписываемся к каналу с того времени, которое указано в бекапе.
