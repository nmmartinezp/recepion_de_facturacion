# Etapa 1: build
FROM golang:1.25-alpine AS builder

# Instalar herramientas necesarias
RUN apk add --no-cache git

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum primero para aprovechar cache de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod tidy

# Copiar todo el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main ./src

# Etapa 2: runtime
FROM alpine:latest

# Crear directorio de trabajo
WORKDIR /app

# Copiar binario compilado
COPY --from=builder /app/main .

# Puerto expuesto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]