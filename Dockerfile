# Используем официальный Go-образ
FROM golang:1.23

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd

# Указываем порт, на котором работает приложение
EXPOSE 8080

# Стартуем приложение
CMD ["./main"]
