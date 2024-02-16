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
- [x] Integração contínua com [GitHub Actions](https://github.com/lucasdanielgeo/customer-history-importer/actions)
- [x] Testes automatizados.

## Instruções de Execução Local

1. **Clone este repositório:**
    ```bash
    git clone https://github.com/lucasdanielgeo/customer-history-importer.git
    ```
2. Navegue até o diretório do projeto:

    ```bash
    cd customer-history-importer
    ```
    **Certifique-se** de que o arquivo **.env** esteja configurado. Deixei um arquivo com os dados para o ambiente de desenvolvimento no repositório. 
    
    Caso queira configurar o seu para ambiente de produção, sugiro não descartar o atual.

3. Execute os seguintes comandos com Docker Compose:
    ```bash
    docker-compose build // apenas na primeira vez
    docker-compose up
    ```
4. Para terminar os containers:
    ```bash
    docker-compose down
    ```
Estes comandos irão construir as imagens Docker e iniciar os contêineres conforme definido no arquivo `docker-compose.yml`.

Uma vez que as imagens tenham sido construídas com sucesso, não será necessário executar novamente o comando `docker-compose build` nas próximas execuções.

## Rodando os testes

Para executar os testes, você pode usar alguns dos comandos disponíveis no Makefile:

- `make tests`: Executa os testes normalmente.
- `make tests-v`: Executa os testes no modo verbose, fornecendo mais informações sobre o processo de teste.
- `make tests-c`: Executa testes, avalia cobertura de testes gerando um relatório html.
