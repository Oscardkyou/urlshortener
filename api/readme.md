Работаем над API движка
API Engine - это HTTP API, используемый клиентом командной строки для связи с демоном. Он также может использоваться сторонним программным обеспечением для управления демоном.

Он состоит из различных компонентов в этом репозитории:

api/swagger.yaml Чванливое определение API.
api/types/ Типы, общие как для клиента, так и для сервера, представляющие различные объекты, опции, ответы и т.д. Большинство из них написаны вручную, но некоторые генерируются автоматически на основе определения Swagger. Смотрите # 27919, как продвигается работа над этим.
cli/ Клиент командной строки.
client/ Клиент Go, используемый клиентом командной строки. Он также может использоваться сторонними программами Go.
daemon/ Демон, который обслуживает API.
Определение чванства
API определяется определением Swagger в api/swagger.yaml. Это определение может быть использовано для:

Автоматически генерирует документацию.
Автоматически генерируют сервер и клиент Go. (Незавершенная работа.)
Предоставить машиночитаемую версию API для анализа того, что он может делать, автоматической генерации клиентов для других языков и т.д.
Обновление документации по API
Документация по API полностью создается на основе api/swagger.yaml. Если вы вносите обновления в API, отредактируйте этот файл, чтобы отразить изменения в документации.

Файл разделен на два основных раздела:

definitions, который определяет объекты многократного использования, используемые в запросах и ответах
paths, который определяет конечные точки API (и некоторые встроенные объекты, которые не обязательно использовать повторно)
Чтобы внести правку, сначала найдите конечную точку, которую вы хотите отредактировать в разделе paths, затем внесите необходимые правки. Конечные точки могут ссылаться на объекты многократного использования с помощью $ref, которые можно найти в definitions разделе.

Надеюсь, в файле достаточно примеров, чтобы вы могли скопировать аналогичный шаблон из других частей файла (например, добавив новые поля или конечные точки), но для получения полной ссылки смотрите спецификацию Swagger.

swagger.yaml проверяетсяhack/validate/swagger, чтобы убедиться, что это правильное определение Swagger. Это полезно при внесении изменений, чтобы убедиться, что вы поступаете правильно.

Просмотр документации по API
При внесении изменений в swagger.yaml вам может потребоваться проверить сгенерированную документацию API, чтобы убедиться, что она отображается правильно.

Запустите, make swagger-docs и будет запущен предварительный просмотр по адресу http://localhost. Некоторые стили могут быть неверными, но вы сможете убедиться, что он генерирует правильную документацию.

Производственная документация генерируется продавцом swagger.yaml в docker/ docker.github.io.