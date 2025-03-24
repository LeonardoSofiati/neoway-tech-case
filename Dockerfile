# Build stage
FROM golang:latest AS builder

WORKDIR /app

# Copia os arquivos do projeto
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Copia o restante do código
COPY . .

# Instala o Swag para gerar a documentação
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Gera a documentação Swagger
RUN swag init --output docs --dir ./cmd/api,./internal/infrastructure/api/handlers,./internal/domain/customer/dto

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go

# Final stage
FROM alpine:latest

# Instala dependências para o timezone
RUN apk --no-cache add tzdata

# Define o timezone
ENV TZ=America/Sao_Paulo

# Copia o binário e a documentação Swagger gerada
COPY --from=builder /app/api /
COPY --from=builder /app/docs /docs

# Copia o script de entrypoint
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Define o entrypoint
ENTRYPOINT ["/entrypoint.sh"]
