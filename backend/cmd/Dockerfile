# Используем официальный образ Go
FROM golang:1.23 AS build

WORKDIR /app
COPY . ./

RUN go mod download
RUN go build -o myapp ./cmd/main.go

# Этап с запуском Go-приложения
FROM gcr.io/distroless/base

# Копируем папку с шаблонами
COPY --from=build /app/myapp /myapp
COPY --from=build /app/templates /templates

EXPOSE 8000
CMD ["/myapp"]
