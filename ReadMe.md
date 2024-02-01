# Тестовое задание для Effective Mobile
Сервис, который получает ФИО, дополняет наиболее вероятным возрастом, полом и национальностью из открытых API и сохраняет в БД postgres.

## Создание персоны

```
POST http://localhost:8080/persons
Content-Type: application/json

{
  "name": "Manya",
  "surname": "Whick",
  "patronymic": "Ivanovna"
}

```
Поле `patronymic` необязательное

## Удаление персоны
```
DELETE http://localhost:8080/persons/14
```

## Получение персон

Для фильтрации используются поля:
* name
* surname
* patronymic

Для пагинации используется поле `page`. Размер каждой страницы - 5 элементов.

```
GET http://localhost:8080/persons?page=1&surname=Sidorov
``` 

## Обновление персоны
Для обновления используются поля:
* name
* surname
* patronymic

Обновляются только переданные поля.

При обновлении поля `name` обновляются поля `age`, `gender`, `nationalize`.
```
PUT http://localhost:8080/persons/16
Content-Type: application/json

{
  "surname": "Sidorov"
}
```
