# Обязательно к прочтению

1) я написал тесты, запускать их надо через bash скрипт (bash test.sh). Они запустятся локально, используют тестововую бд (в ней должны быть и posts и comments таблицы)
2) проблему n + 1 я решил - в бд запрос по получению комментариев делается рекурсиво. Для in-memory хранилища я с сделал две мапы. Comments - ключ это id, значение слайс с комментиями под этим комментарием/постом. Если нет parentID, то прикрепляем к id поста, иначе прикреляем к parentID. Когда захотим вывести комментарии - будем добавлять их в результирующий слайс рекурсивно.
3) Реализовал выбор хранилища - в докерфайле 13 строка

```Dockerfile
CMD ["./server"]
```

тогда будет in-memory хранилище

```Dockerfile
CMD ["./server", "db"]
```

при смене хранилища надо пересобрать

```bash
docker-compose up --build
```

4) я не сделал subsriptions

## Как запустить?

я выставил для in-memory, поменять - читать предыдущий пункт

```bash
docker-compose up
```

Query и mutations лежат в helper. Оттуда скопируйте все, вставьте в GraphyQL (это в бразуере).
P S порядок такой - делаем пост, коммент, коммент1, коммент2. 
Дальше по вашему усмотрению