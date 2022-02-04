Команды для запуска сервиса:

  ` docker build -t mygrpc . `
 
  `  docker-compose up `

Команды для отправки запросов через grpc клиент:

   создание пользователя:
  
    ` client/client create <user_login> <user_email>`

  удаление пользователя:
  
    ` client/client delete <user_id>`
  
  получить список последних 1000 пользователей:
  
    ` client/client list `
