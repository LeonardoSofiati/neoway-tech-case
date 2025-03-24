# 🚀 API Neoway - Case Técnico

Este projeto é uma API desenvolvida em **Golang** com suporte a **Swagger** e banco de dados **PostgreSQL**. O ambiente é gerenciado usando **Docker** e **Docker Compose**, garantindo fácil configuração e execução.

## 📌 Tecnologias Utilizadas
- **Golang** (backend)
- **PostgreSQL** (banco de dados)
- **Docker** e **Docker Compose** (gerenciamento de ambiente)
- **Swagger** (documentação da API)

## 📂 Estrutura do Projeto
```
📦 projeto
├── cmd/
│   ├── api/
│   │   └── main.go  # Arquivo principal da API
├── internal/
│   ├── domain/
│   │   ├── customer/
│   │   │   ├── dto/         # Definição dos DTOs
│   │   │   ├── entity/      # Entidades do domínio
│   │   │   ├── repository/  # Repositórios do domínio
│   │   │   └── service/     # Lógica de serviço do domínio
│   │   ├── shared/
│   │   │   ├── entity/      # Entidades compartilhadas
│   │   │   └── repository/  # Repositórios compartilhados
│   ├── infrastructure/
│   │   ├── api/
│   │   │   └── handlers/     # Handlers das rotas da API
│   │   ├── database/
│   │   │   ├── config/       # Configurações de banco de dados
│   │   │   └── repository/   # Repositórios do banco de dados
│   ├── internal-errors/      # Gerenciamento de erros internos
│   │   ├── error.go          # Definição de tipos e mensagens de erro
│   │   └── handler.go        # Handler de erros
│   ├── usecase/
│   │   └── customer/         # Casos de uso específicos para customer
│   │       ├── create/       # Caso de uso para criação de customer
│   │       ├── delete/       # Caso de uso para exclusão de customer
│   │       ├── find/         # Caso de uso para busca de customer
│   │       └── list/         # Caso de uso para listar customers
├── docs/  # Documentação gerada pelo Swagger
├── Dockerfile  # Configuração do container
├── docker-compose.yml  # Configuração do ambiente
├── go.mod  # Dependências do Go
├── go.sum  # Checksum das dependências
├── Makefile  # Comandos úteis
├── entrypoint.sh  # Script de inicialização
└── README.md  # Documentação do projeto
```

## 🔧 Pré-requisitos
Antes de começar, você precisará ter instalado na sua máquina:
- **Docker** e **Docker Compose**
- **Golang** (caso queira rodar sem Docker)

## 🚀 Como Rodar o Projeto

### 1️⃣ Clonar o repositório
```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```

### 2️⃣ Subir os containers com Docker Compose
```bash
docker compose -f docker-compose.yml up --build
```
Isso irá:
- Criar e iniciar um banco de dados PostgreSQL.
- Construir a imagem da API e rodá-la.
- Gerar automaticamente a documentação Swagger.

### 3️⃣ Acessar a API
A API estará rodando em `http://localhost:8080`

### 4️⃣ Acessar a Documentação Swagger com todos os edpoints e a possibilidade de testar sem necessidade do POSTMAN
Abra no navegador:
```
http://localhost:8080/swagger/
```

## 🔄 Rodando sem Docker (Opcional)
Se quiser rodar localmente sem Docker:

### 1️⃣ Subir o container em PostgreSQL
```bash
docker run --name postgres_neoway \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_USER=neoway_dev \
  -e POSTGRES_DB=neoway_dev \
  -p 5432:5432 \
  -d postgres
```

### 2️⃣ Instalar as dependências do projeto
```bash
go mod tidy
go mod download
```

### 3️⃣ Gerar a documentação Swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init --output docs --dir ./cmd/api,./internal/infrastructure/api/handlers,./internal/domain/customer/dto
```

### 4️⃣ Executar a API
```bash
go run cmd/api/main.go
```

---
Desenvolvido por [Leonardo Sofiati Buscariolo](https://github.com/seu-usuario) 🚀

