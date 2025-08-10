Запуск через docker-compose.

Если что, подключиться к базе можно с хостовой машины:
* Порт: 54320
* POSTGRES_DB: order_service_db
* POSTGRES_USER: myuser
* POSTGRES_PASSWORD: mypassword

Снаружи подключиться к Kafka по порту 9094.

Включение/выключение отправки тестовых данных по переменной PRODUCER_ENABLED в docker-compose.

Включение / выключение получения тестовых данных - CONSUMER_ENABLED.


Документация на API: http://localhost:8080/swagger/index.html
