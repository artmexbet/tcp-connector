# TCP Connector

Простая утилита для проверки TCP-портов, написанная на Go.

## Возможности

- Проверка доступности TCP-порта по указанному IP.
- REST API интерфейс.
- Graceful shutdown (плавное завершение работы).
- Конфигурация через переменные окружения.

## Требования

- Go 1.18+

## Установка

```bash
git clone https://github.com/yourusername/tcp-conntector.git
cd tcp-conntector
go mod download
```

## Использование

### Запуск сервера

```bash
# Порт по умолчанию 8080
go run cmd/app/main.go

# Или укажите свой порт
PORT=9090 go run cmd/app/main.go
```

### Проверка порта

Отправьте POST запрос на `/check`:

```bash
curl -X POST http://localhost:8080/check \
  -H "Content-Type: application/json" \
  -d '{"ip": "google.com", "port": 80}'
```

Ответ:
```json
{
  "ip": "google.com",
  "port": 80,
  "status": "open"
}
```

## Спецификация API

Полную документацию API смотрите в файле [openapi.yaml](openapi.yaml).

## Решения по реализации
Использовал везде только стандартную библиотеку Go, 
чтобы минимизировать зависимости и упростить развертывание.

По-хорошему, все константы вынести бы в конфиги, но для 
простоты примера оставил их в коде.