# Cборка
## Этап 1
Для сборки вам понадобиться Golang версии 1.22.х

## Этап 2
Запуск скрипта для сборки:
```sh
./build.sh
```

> Если у вас ошибка прав запустить команду: sudo chmod 755 build.sh

# Запуск
для того что бы запустить проект после сборки вам небходимо запуститить собранный файл командой:
```sh
./memory_wall
```

# Документация API
### Пути парсинга данных
<details>
<summary><code>POST</code> <code><b>/api/parse/docx</b></code> <code>(Возвращает содержимое файла в формате json)</code></summary>

#### Параметры
>| Name | Type   | Data type |Description|
>|------|--------| ----------|-----------|
>|files |required| array     |N/A        |

#### Ответ
>|Code| Content-Type   | response |
>|----|----------------|----------|
>|200 |application/json|```json```|

##### Структура ответа JSON
```json
{
    "data":[
        {
            "filename": "string",
            "human_info": {
                "name": "string",
                 "first_name": "string",
                "last_name": "string",
                "middle_name": "string",
                "description": "string",
                "birthday": "date|string",
                "deathday": "date|string",
                "place_of_birth": "string",
                "date_and_place_of_conscription": "string",
                "military_rank_and_position": "string",
                "awards": "[]string",
                "images": "[]byte"
            }
        }
    ]
}
```
</details>

