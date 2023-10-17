# GoUrlShortener

A simple URL shortener written in Go.


## Руководство по подключению Go приложения к PostgreSQL

1. **Установка PostgreSQL**
   - Если у вас еще не установлен PostgreSQL, установите его с [официального сайта](https://www.postgresql.org/download/).
   - После установки убедитесь, что PostgreSQL сервер запущен и вы можете подключиться к нему с помощью утилиты `psql` или другого клиента.

2. **Создание базы данных и пользователя**
   - Запустите `psql` и выполните следующие команды:
     ```sql
     CREATE DATABASE your_db_name;
     CREATE USER your_user_name WITH PASSWORD 'your_password';
     GRANT ALL PRIVILEGES ON DATABASE your_db_name TO your_user_name;
     ```

3. **Подключение к PostgreSQL из Go**
   - Добавьте необходимые зависимости в ваш проект:
     ```bash
     go get -u github.com/lib/pq
     ```
   - В вашем коде импортируйте пакет `_ "github.com/lib/pq"`.
   - Используйте строку подключения следующего формата для создания подключения:
     ```go
     connStr := "user=your_user_name password=your_password dbname=your_db_name sslmode=disable"
     db, err := sql.Open("postgres", connStr)
     ```

4. **Миграции (опционально)**
   - Для управления структурой вашей базы данных вы можете рассмотреть использование инструмента миграций, например, [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

5. **Обработка ошибок**
   - После каждой операции с базой данных (создание подключения, выполнение запроса и т. д.) всегда проверяйте наличие ошибок и обрабатывайте их соответствующим образом.

6. **Закрытие подключения**
   - После завершения работы с базой данных убедитесь, что вы закрыли подключение с помощью метода `db.Close()`.

7. **Безопасность**
   - Никогда не храните учетные данные для подключения в открытом виде в вашем коде. Используйте переменные окружения или конфигурационные файлы.
   - Если возможно, настройте ваш сервер PostgreSQL так, чтобы он принимал соединения только с определенных IP-адресов или диапазонов IP-адресов.
