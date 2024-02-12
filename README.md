# Serviço de Manipulação e Persistência de Dados

## Objetivo
Desenvolver um serviço para manipulação de dados e persistência de dados de histórico de compras de usuários em um banco de dados relacional.

## Requisitos Obrigatórios
- Criar um serviço utilizando Go, Python ou Javascript/Typescript.
- Receber um arquivo CSV ou TXT como entrada.
- Persistir os dados no banco de dados PostgreSQL.
- Fazer o split dos dados em colunas no banco de dados.
- Utilizar Docker Compose para a execução do projeto.
- Realizar higienização dos dados após a persistência.
- Validar os CPFs/CNPJs contidos nos dados.

## Desejável

- [x] Execução direta no serviço da linguagem escolhida ou em SQL.
- [] Integração contínua com Github Actions
- [] Testes automatizados.

## Instruções de Execução Local
1. Clone este repositório: `git clone github.com/lucasdanielgeo/customer-history-importer.git`
2. Navegue até o diretório do projeto: `cd nome_do_projeto`
3. Execute o Docker Compose: `docker-compose up`
