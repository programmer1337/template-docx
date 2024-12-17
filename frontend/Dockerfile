FROM node:22 AS build-app
WORKDIR /app

RUN apt update && apt install git autoconf automake gcc make g++ libtool -y
RUN git init

# Копируем package.json и package-lock.json
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем остальные файлы проекта
COPY . ./

# Пишем переменную в .env
# ARG ENV_ARG=$ENV_ARG
# RUN echo -e "$ENV_ARG" >> ./.env

ARG API_URL=localhost
ENV API_URL=${API_URL}

# Строим приложение для продакшен-режима
RUN npm run build

# Этап 2: Запуск с использованием Nginx
FROM nginx:alpine

# Копируем скомпилированные файлы из стадии сборки в Nginx
COPY --from=build-app /app/build /usr/share/nginx/html

# Экспонируем порт 80
EXPOSE 80

# Запускаем Nginx для обслуживания приложения
CMD ["nginx", "-g", "daemon off;"]
