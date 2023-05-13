### Тестовая задача: необходимо создать консольное приложение-сервис, которое принимает HTTP POST запросы:

1. **По пути POST `/redis/incr` с json вида**

    {
    "key": "age",
    "value": 19
    
    },
    подключается к redis (хост и порт указываются при запуске в параметрах `-host` и `-port`),
    инкрементирует значение по ключу, указанному в "key" на значение из value", и
    возвращает его в виде:
    
    {
    "value": 20
    
    }

2. **По пути POST `/sign/hmacsha512` с json вида**
    
    {
    "text": "test",
    "key": "test123"
    
    }
    
    и возвращает HMAC-SHA512 подпись значения из "text" по ключу "key" в виде hex строки

3. **По пути POST `/postgres/users` с json вида**

    {
    "name": "Alex",
    "age": 21
    
    }
    
    создает в базе данных postgres таблицу users, если она не существует, добавляет в нее
    строку (“Alex”, 21) и возвращает идентификатор добавленного пользователя в виде
    {
    
    “id”: 1
    }

**Дополнительные условия:**

1. Можно использовать любые библиотеки для работы с http, redis и postgres;
2. Приложение должно быть протестировано unit-тестами (любой тестовый фреймворк);
3. Наибольшее внимание следует уделить качеству коду.

**Результатом этой задачи должен стать проект на гите.**

### На что будем обращать внимание

- погруженность в тему;
- процесс поиска решения;
- качество кода.
