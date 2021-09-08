Тестовое задание на позицию стажера-бекендера
==============================================

Задание доступно по [ссылке](https://github.com/olteffe/avitochat/blob/master/task.md)

Примененные технологии
-----------------------

1. REST API.
2. Golang.
3. ORM.
4. Чистая архитектура.
5. Шаблон [Go Project Layout](https://github.com/golang-standards/project-layout).
8. [Кодогенерация](https://github.com/openapitools/openapi-generator) на основе OAS3.0.
6. Redis.
7. PostgreSQL.
9. Docker.

Зависимости
------------

Использованы следующие внешние библиотеки:
* [Echo](https://github.com/labstack/echo) - web framework.
* [Gorm](https://github.com/go-gorm/gorm) - ORM для Golang.
* [Google/uuid](https://github.com/google/uuid) - генерирует и проверяет UUID на основе RFC4122 и DCE 1.1.
* [Mockery](https://github.com/vektra/mockery) - генератор кода для имитации интерфейсов Golang.
* [Testify](https://github.com/stretchr/testify/) - framework для написания тестов на Golang.

Документация
--------------

Документация представлена в виде OpenAPI [спецификации](https://github.com/olteffe/avitochat/blob/master/api/openapi.yaml).

![db here](https://github.com/olteffe/avitochat/blob/master/assets/oapi.png)

ER-диаграмма
--------------
![db here](https://github.com/olteffe/avitochat/blob/master/assets/db_avitochat.png)


Архитектура приложения
-----------------------

![alt-текст](https://github.com/olteffe/avitochat/blob/master/assets/arch.png "Архитектура приложения")

Этапы разработки
-----------------

1. Анализ ТЗ.
2. Разрабатываем архитектуру(определяемся со стеком технологий). Рисуем диаграммы: бд, приложения
3. Создаем спецификацию OAS.
4. Генерируем шаблон сервера.
5. Программируем.
6. Пишем тесты.
7. Разворачиваем приложение в дев среде(docker контейнеры).
8. Исправляем баги
9. Используем танк/ApacheBench.
