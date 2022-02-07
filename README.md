Команды для запуска сервиса:
  
  `  docker-compose up -d `

Компиляция клиента grpc:
  
   ` cd client ` \
   ` go build client.go `

Команды для отправки запросов через grpc клиент:

   создание пользователя:
  
    ` client/client create <user_login> <user_email> `

  удаление пользователя:
  
    ` client/client delete <user_id> `
  
  получить список последних 1000 пользователей:
  
    ` client/client list `
