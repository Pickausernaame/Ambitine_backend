# Ambitine_backend
Бекэнд для невероятно мятного приложения Ambitine .

Команда для запуска контейнеров:
```
sudo docker-compose up --build
```

Команда для локального подключения к контейнеру бд, который крутится на вашей тачке:
```
psql -U ambiuser -h 0.0.0.0 -p 54321 --dbname=ambitine
```


Команда для включения citext'а вручную:
```
 psql -U ambiuser -h 0.0.0.0 -p 54321 --dbname=ambitine --echo-all --command 'create extension if not exists "citext";'
```
