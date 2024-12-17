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

---

## Как использовать

1. Запустите исполняемый файл `main.exe`.
2. Отправьте POST-запрос (например, через cURL) на URL: `localhost:8080/api/v1/calculate`.
3. Получите ответ.

---

## Примеры использования

### Успешный запрос

**Пример запроса через cURL:**

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2*2+2"
}'
```

**Пример ответа:**

```json
{
    "result": "6"
}
```

**HTTP-код ответа:** `200` (всё прошло успешно).

---

### Ошибка 422 (Некорректное выражение)

**Пример запроса через cURL:**

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "0.25+0.251++"
}'
```

**Пример ответа:**

```json
{
    "error": "Expression is not valid"
}
```

**HTTP-код ответа:** `422` (Unprocessable Entity).

---

### Ошибка 500 (Отправка не POST-запроса)

**Пример запроса через cURL:**

```bash
curl --location --request GET 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "0.25+0.251++"
}'
```

**Пример ответа:**

```json
{
    "error": "Internal server error"
}
```

**HTTP-код ответа:** `500` (Internal Server Error).

---

## Примечания

- Убедитесь, что порт `8080` не занят другими приложениями.
- Поддерживаются стандартные арифметические операции: сложение, вычитание, умножение, деление и скобки для изменения порядка операций.
- Поддерживаются только POST-запросы, отправка других приведёт к ошибке.

