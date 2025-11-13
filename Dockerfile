# Etapa 1: build
FROM golang:1.25-alpine AS builder

# Instalar herramientas necesarias
RUN apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum primero para aprovechar cache de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod tidy

# Copiar todo el código fuente
COPY . .

# Generar documentación Swagger
RUN swag init -g src/main.go -o src/docs

# Compilar la aplicación
RUN go build -o main ./src

# Etapa 2: runtime
FROM alpine:latest

# Crear directorio de trabajo
WORKDIR /app

# Copiar binario compilado
COPY --from=builder /app/main .

# Puerto expuesto
EXPOSE 3002

# Comando para ejecutar la aplicación
CMD ["./main"]