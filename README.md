# Сервис подсчёта арифметических выражений

Данная программа представляет собой API-интерфейс, реализованный на Go, который позволяет отправлять POST-запросы с телом:

```json
{
    "expression": "выражение"
}
```

и получать ответ вида:

```json
{
    "result": "результат"
}
```

## Как использовать

1. Запустите исполняемый файл `main.exe`.
2. Отправьте POST-запрос (например, через cURL) на URL: `localhost:8080/api/v1/calculate`.
3. Получите ответ.

## Примеры использования

### Пример запроса через cURL

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2*2+2"
}'
```

### Пример ответа

```json
{
    "result": "6"
}
```

HTTP-код ответа: `200` (всё прошло успешно).

### Примеры использования(некоректное выражение)
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "0.25+0.251++"
}
'
```
### Ответ:
```json
{
    "error": "Expression is not valid"
}
```
HTTP код ответа: 422

### Примеры использования(отправка не POST запроса)
```bash
curl --location --request GET 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "0.25+0.251++"
}
'
```

### Ответ
```bash
{
    "error": "Internal server error"
}
```

## Примечания

- Убедитесь, что порт `8080` не занят другими приложениями.
- Поддерживаются стандартные арифметические операции: сложение, вычитание, умножение, деление и скобки для изменения порядка операций.
- Поддерживаются только POST запросы, отправка других приведёт к ошибке
