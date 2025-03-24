#!/bin/sh

# Espera o banco de dados ficar pronto antes de iniciar a API
echo "Aguardando o banco de dados iniciar..."

while ! nc -z db 5432; do
    sleep 2
done

echo "Banco de dados disponível. Iniciando a API..."

# Inicia a aplicação
exec ./api
