# go-gateway-api

Este projeto é uma API Gateway desenvolvida em Go.

## Estrutura do Projeto

- `cmd/app/`: Contém o código-fonte principal da aplicação.
- `internal/`: Inclui pacotes internos utilizados pela aplicação.
- `migrations/`: Arquivos de migração para o banco de dados.
- `.env.example`: Exemplo de arquivo de variáveis de ambiente.
- `.gitignore`: Lista de arquivos e diretórios ignorados pelo Git.
- `docker-compose.yaml`: Configuração para execução com Docker Compose.
- `go.mod`: Gerenciamento de dependências do Go.
- `test.http`: Arquivo para testes de requisições HTTP.

## Pré-requisitos

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started) e [Docker Compose](https://docs.docker.com/compose/install/)

## Configuração

1. Copie o arquivo `.env.example` para `.env` e ajuste as variáveis conforme necessário.
2. Execute `docker-compose up` para iniciar os serviços.

## Uso

Após iniciar os serviços, a API estará disponível em `http://localhost:PORTA`. Utilize ferramentas como [Postman](https://www.postman.com/) ou `curl` para interagir com a API.

## Testes

Utilize o arquivo `test.http` para realizar testes de requisições HTTP.

## Contribuição

Contribuições são bem-vindas! Siga os passos abaixo:

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas alterações (`git commit -m 'Adiciona nova feature'`).
4. Faça o push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
